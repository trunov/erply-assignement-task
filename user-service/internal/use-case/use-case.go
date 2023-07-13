package use_case

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/trunov/erply-assignement-task/user-service/internal/domain"
	"github.com/trunov/erply-assignement-task/user-service/internal/repository/erply"
)

type Storage interface {
	GetCustomer(ctx context.Context, id int) (*erply.Customer, error)
	StoreCustomer(ctx context.Context, customer *erply.Customer) error
}

type Erply interface {
	GetCustomer(ctx context.Context, sessionKey, clientCode, customerID string) (erply.Customer, error)
	AddCustomer(ctx context.Context, sessionKey, clientCode string, customer domain.CustomerInput) error
	ErplyAuthentication(ctx context.Context, clientCode, username, password string) (erply.GetVerifyUserResponse, error)
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

func (c *useCase) GetCustomer(ctx context.Context, sessionKey, clientCode, customerID string) (erply.Customer, error) {
	num, err := strconv.Atoi(customerID)
	if err != nil {
		return erply.Customer{}, err
	}

	storageCustomer, err := c.storage.GetCustomer(ctx, num)
	if err != nil && err != sql.ErrNoRows {
		return erply.Customer{}, err
	}

	if storageCustomer != nil {
		return *storageCustomer, nil
	}

	customer, err := c.erply.GetCustomer(ctx, sessionKey, clientCode, customerID)
	if err != nil {
		return customer, err
	}

	err = c.storage.StoreCustomer(ctx, &customer)
	if err != nil {
		return erply.Customer{}, err
	}

	return customer, nil
}

func (c *useCase) AddCustomer(ctx context.Context, sessionKey, clientCode string, newCustomer domain.CustomerInput) error {
	err := c.erply.AddCustomer(ctx, sessionKey, clientCode, newCustomer)
	if err != nil {
		return err
	}

	return nil
}
