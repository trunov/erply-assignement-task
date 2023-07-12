package erply

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type GetVerifyUserResponse struct {
	Status  Status     `json:"status"`
	Records []UserInfo `json:"records"`
}

type Status struct {
	Request           string  `json:"request"`
	RequestUnixTime   int64   `json:"requestUnixTime"`
	ResponseStatus    string  `json:"responseStatus"`
	ErrorCode         int     `json:"errorCode"`
	GenerationTime    float64 `json:"generationTime"`
	RecordsTotal      int     `json:"recordsTotal"`
	RecordsInResponse int     `json:"recordsInResponse"`
}

type UserInfo struct {
	UserID       string `json:"UserID"`
	EmployeeName string `json:"employeeName"`
	SessionKey   string `json:"sessionKey"`
}

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
