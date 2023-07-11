package erply

import (
	"context"
	"fmt"
	"net/http"

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

func (c *client) GetCustomer(ctx context.Context, id int) {

}
