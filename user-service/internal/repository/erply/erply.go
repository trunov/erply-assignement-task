package erply

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/trunov/erply-assignement-task/user-service/internal/domain"
)

var ErrRequestFailed = errors.New("request did not make through")

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

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormDataFromValues(payload).
		Post(requestURL)

	if err != nil {
		return Customer{}, err
	}

	err = json.Unmarshal(resp.Body(), &response)
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

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormDataFromValues(payload).
		Post(requestURL)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(resp.Body(), &response)
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

	var response SaveCustomerResponse

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormDataFromValues(payload).
		Post(requestURL)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if response.Status.ResponseStatus != "ok" {
		return ErrRequestFailed
	}

	return nil
}
