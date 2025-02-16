// Code generated by goa v3.19.1, DO NOT EDIT.
//
// items client
//
// Command:
// $ goa gen github.com/kormiltsev/be/internal/api/design -o ./internal/api

package items

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "items" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	CreateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
}

// NewClient initializes a "items" service client given the endpoints.
func NewClient(list, create, delete_ goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		CreateEndpoint: create,
		DeleteEndpoint: delete_,
	}
}

// List calls the "list" endpoint of the "items" service.
// List may return the following errors:
//   - "not_found" (type *NotFoundError): is a common error response for not found
//   - "bad_request" (type *BadRequestError): is a common error response for bad request
//   - "internal" (type *InternalError): is a common error response for internal error
//   - error: internal error
func (c *Client) List(ctx context.Context) (res *RenamemeItems, err error) {
	var ires any
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*RenamemeItems), nil
}

// Create calls the "create" endpoint of the "items" service.
// Create may return the following errors:
//   - "not_found" (type *NotFoundError): is a common error response for not found
//   - "bad_request" (type *BadRequestError): is a common error response for bad request
//   - "internal" (type *InternalError): is a common error response for internal error
//   - error: internal error
func (c *Client) Create(ctx context.Context, p *CreatePayload) (err error) {
	_, err = c.CreateEndpoint(ctx, p)
	return
}

// Delete calls the "delete" endpoint of the "items" service.
// Delete may return the following errors:
//   - "not_found" (type *NotFoundError): is a common error response for not found
//   - "bad_request" (type *BadRequestError): is a common error response for bad request
//   - "internal" (type *InternalError): is a common error response for internal error
//   - error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}
