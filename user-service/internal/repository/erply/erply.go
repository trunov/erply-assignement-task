package erply

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/trunov/erply-assignement-task/user-service/internal/domain"
)

type client struct {
	client *resty.Client
	addr   string
}

func New(hc *http.Client, clientCode string) *client {
	addr := fmt.Sprintf("https://%s.erply.com/api/", clientCode)

	return &client{
		client: resty.NewWithClient(hc),
		addr:   addr,
	}
}

func (c *client) GetCustomer(ctx context.Context, sessionKey, clientCode, customerID string) (Customer, error) {
	requestURL := c.addr
	payload := url.Values{
		"sessionKey":      {sessionKey},
		"request":         {"getCustomers"},
		"customerID":      {customerID},
		"sendContentType": {"1"},
		"clientCode":      {clientCode},
	}

	var response GetCustomerResponse

	_, err := c.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormDataFromValues(payload).
		SetResult(&response).
		Post(requestURL)

	if err != nil {
		return Customer{}, err
	}

	customer := response.Records[0]

	return customer, nil
}

func (c *client) ErplyAuthentication(ctx context.Context, clientCode, username, password string) (GetVerifyUserResponse, error) {
	requestURL := c.addr
	payload := url.Values{
		"username":        {username},
		"password":        {password},
		"request":         {"verifyUser"},
		"sendContentType": {"1"},
		"clientCode":      {clientCode},
	}

	var response GetVerifyUserResponse

	_, err := c.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormDataFromValues(payload).
		SetResult(&response).
		Post(requestURL)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *client) AddCustomer(ctx context.Context, sessionKey, clientCode string, customer domain.CustomerInput) error {
	requestURL := c.addr
	payload := url.Values{
		"sessionKey":      {sessionKey},
		"request":         {"saveCustomer"},
		"clientCode":      {clientCode},
		"firstName":       {customer.FirstName},
		"lastName":        {customer.LastName},
		"email":           {customer.Email},
		"phone":           {customer.Phone},
		"twitterID":       {customer.TwitterID},
		"sendContentType": {"1"},
	}

	_, err := c.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormDataFromValues(payload).
		Post(requestURL)

	if err != nil {
		return err
	}

	return nil
}
