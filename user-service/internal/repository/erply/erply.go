package erply

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
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

func (c *client) GetCustomer(ctx context.Context, sessionKey, clientCode, customerID string) (*resty.Response, error) {
	requestURL := c.addr
	payload := url.Values{
		"sessionKey":      {sessionKey},
		"request":         {"getCustomers"},
		"customerID":      {customerID},
		"sendContentType": {"1"},
		"clientCode":      {clientCode},
	}

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormDataFromValues(payload).
		Post(requestURL)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
