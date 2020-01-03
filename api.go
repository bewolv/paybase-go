package paybase

import (
	"fmt"
	"github.com/bewolv/paybase-go/account"
	"github.com/bewolv/paybase-go/auth"
	"github.com/bewolv/paybase-go/customer"

	"context"
)

func (api *Client) GetStatus(ctx context.Context, outp interface{}) error {
	return api.getMethod(ctx, "/v1/status", nil, outp)
}

func (api *Client) CreateIncorporatedBusiness(ctx context.Context, body *customer.CreateIncorporatedBusiness, response *customer.IncorporatedBusiness) error {
	return api.postMethod(ctx, "POST", "/v1/customers/incorporated-business", body, response)
}

func (api *Client) CreateCustomerAuthenticationToken(ctx context.Context, body *auth.CreateToken, response *auth.AuthenticationToken) error {
	path := fmt.Sprintf(`/v1/%s/tokens`, body.ID)
	return api.postMethod(ctx, "POST", path, body, response)
}

func (api *Client) CreateAccount(ctx context.Context, body *account.Create, response *account.Account) error {
	path := "/v1/accounts"
	return api.postMethod(ctx, "POST", path, body, response)
}

func (api *Client) AddBusinessPerson(ctx context.Context, body *customer.CreateBusinessPerson, response *customer.BusinessPerson) error {
	path := fmt.Sprintf("/v1/customer/%s/persons", body.CustomerID)
	return api.postMethod(ctx, "POST", path, body, response)
}

func (api *Client) UpdateIncorporatedBusiness(ctx context.Context, body *customer.UpdateIncorporatedBusiness, response *customer.IncorporatedBusiness) error {
	path := fmt.Sprintf("/v1/customer/incorporated-business/%s", body.ID)
	return api.postMethod(ctx, "PATCH", path, body, response)
}
