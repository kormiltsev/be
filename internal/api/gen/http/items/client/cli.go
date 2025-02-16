// Code generated by goa v3.19.1, DO NOT EDIT.
//
// items HTTP client CLI support package
//
// Command:
// $ goa gen github.com/kormiltsev/be/internal/api/design -o ./internal/api

package client

import (
	"encoding/json"
	"fmt"

	items "github.com/kormiltsev/be/internal/api/gen/items"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePayload builds the payload for the items create endpoint from CLI
// flags.
func BuildCreatePayload(itemsCreateBody string) (*items.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(itemsCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"item\": {\n         \"description\": \"Modi ratione assumenda.\",\n         \"icon\": \"Est voluptas quisquam cupiditate.\",\n         \"id\": \"Voluptatibus explicabo deleniti aut quos.\",\n         \"labels\": [\n            \"Ea magnam occaecati illo reprehenderit.\",\n            \"Ipsa doloremque voluptatem voluptatum provident.\"\n         ],\n         \"name\": \"dy\"\n      }\n   }'")
		}
		if body.Item == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("item", "body"))
		}
		if body.Item != nil {
			if err2 := ValidateApplicationVndRenamemeItemJSONRequestBody(body.Item); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	v := &items.CreatePayload{}
	if body.Item != nil {
		v.Item = marshalApplicationVndRenamemeItemJSONRequestBodyToItemsApplicationVndRenamemeItemJSON(body.Item)
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the items delete endpoint from CLI
// flags.
func BuildDeletePayload(itemsDeleteID string) (*items.DeletePayload, error) {
	var id string
	{
		id = itemsDeleteID
	}
	v := &items.DeletePayload{}
	v.ID = id

	return v, nil
}
