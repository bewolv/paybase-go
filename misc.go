package paybase

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// StatusCodeError represents an http response error.
// type httpStatusCode interface { HTTPStatusCode() int } to handle it.
type statusCodeError struct {
	Code   int
	Status string
}

func (t statusCodeError) Error() string {
	return fmt.Sprintf("Paybase error: %s", t.Status)
}

func (t statusCodeError) HTTPStatusCode() int {
	return t.Code
}

func (t statusCodeError) Retryable() bool {
	if t.Code >= 500 || t.Code == http.StatusTooManyRequests {
		return true
	}
	return false
}

func checkStatusCode(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {

		// TODO implement logging
		return statusCodeError{Code: resp.StatusCode, Status: resp.Status}
	}

	return nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func getResource(ctx context.Context, api *Client, endpoint string, valuesStruct interface{}, intf interface{}) error {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	if valuesStruct != nil {
		values, err := Values(valuesStruct)
		if err != nil {
			return err
		}
		req.URL.RawQuery = values.Encode()
	}

	req.Header.Set("Content-Type", "application/json")

	return doPost(ctx, api, req, newJSONParser(intf))
}

type responseParser func(*http.Response) error

func newJSONParser(dst interface{}) responseParser {
	return func(resp *http.Response) error {
		return json.NewDecoder(resp.Body).Decode(dst)
	}
}

// post a url encoded form.
func postForm(ctx context.Context, api *Client, method string, endpoint string, body interface{}, outp interface{}) error {
	switch method {
	case "POST":
	case "PATCH":
	case "DELETE":
	default:
		return errors.New("Invalid method")
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req, err := http.NewRequest(method, endpoint, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	return doPost(ctx, api, req, newJSONParser(outp))
}

func doPost(ctx context.Context, api *Client, req *http.Request, parser responseParser) error {
	req = req.WithContext(ctx)
	req.Header.Set("x-token", api.token)

	resp, err := api.httpclient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO implement debugging logic
	if false {
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(string(body))
	}
	err = checkStatusCode(resp)
	if err != nil {
		return err
	}

	return parser(resp)
}
