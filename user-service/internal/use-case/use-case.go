package use_case

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Storage interface {
	GetCustomer(ctx context.Context, id int)
}

type Erply interface {
	GetCustomer(ctx context.Context, sessionKey, clientCode, customerID string) (*resty.Response, error)
}

type useCase struct {
	storage Storage
	erply   Erply
}

func New(storage Storage, erply Erply) *useCase {
	return &useCase{
		storage: storage,
		erply:   erply,
	}
}

func (c *useCase) GetCustomer(ctx context.Context, sessionKey, clientCode, customerID string) (*resty.Response, error) {
	resp, err := c.erply.GetCustomer(ctx, sessionKey, clientCode, customerID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
