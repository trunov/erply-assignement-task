package use_case

import "context"

type Storage interface {
	GetCustomer(ctx context.Context, id int)
}

type Erply interface {
	GetCustomer(ctx context.Context, userID int)
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

func (c *useCase) GetCustomer(ctx context.Context, id int) {

}
