package paybase

import (
	"context"
	"net/http"
	"net/url"
)

const (
	// APIURL paybase.
	SANDAPIURL = "https://api-json.sandbox.paybase.io"
	PRODAPIURL = "https://api-json.paybase.io"
)

// httpClient defines the minimal interface needed for an http.Client to be implemented.
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client for the paybase api.
type ParamOption func(*url.Values)

type Client struct {
	token      string
	endpoint   string
	httpclient httpClient
}

// New builds a paybse client from the provided token and options.
func New(token string, prod bool) *Client {
	s := &Client{
		token:      token,
		httpclient: &http.Client{},
	}
	if prod {
		s.endpoint = PRODAPIURL
	} else {
		s.endpoint = SANDAPIURL
	}

	return s
}

// post to a paybase method.
func (api *Client) postMethod(ctx context.Context, method string, path string, inp interface{}, outp interface{}) error {
	return postForm(ctx, api, method, api.endpoint+path, inp, outp)
}

// get a paybase method.
func (api *Client) getMethod(ctx context.Context, path string, valueStruct interface{}, intf interface{}) error {
	return getResource(ctx, api, api.endpoint+path, valueStruct, intf)
}
