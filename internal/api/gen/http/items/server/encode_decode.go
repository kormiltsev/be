// Code generated by goa v3.19.1, DO NOT EDIT.
//
// items HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/kormiltsev/be/internal/api/design -o ./internal/api

package server

import (
	"context"
	"errors"
	"io"
	"net/http"

	items "github.com/kormiltsev/be/internal/api/gen/items"
	itemsviews "github.com/kormiltsev/be/internal/api/gen/items/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the items
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*itemsviews.RenamemeItems)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewListResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeListError returns an encoder for errors returned by the list items
// endpoint.
func EncodeListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *items.BadRequestError
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "internal":
			var res *items.InternalError
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListInternalResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateResponse returns an encoder for responses returned by the items
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the items create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreatePayload(&body)

		return payload, nil
	}
}

// EncodeCreateError returns an encoder for errors returned by the create items
// endpoint.
func EncodeCreateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *items.BadRequestError
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "internal":
			var res *items.InternalError
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateInternalResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the items
// delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the items delete
// endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewDeletePayload(id)

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete items
// endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "bad_request":
			var res *items.BadRequestError
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "internal":
			var res *items.InternalError
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteInternalResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *items.NotFoundError
			errors.As(v, &res)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalItemsviewsApplicationVndRenamemeItemJSONViewToApplicationVndRenamemeItemJSONResponseBody
// builds a value of type *ApplicationVndRenamemeItemJSONResponseBody from a
// value of type *itemsviews.ApplicationVndRenamemeItemJSONView.
func marshalItemsviewsApplicationVndRenamemeItemJSONViewToApplicationVndRenamemeItemJSONResponseBody(v *itemsviews.ApplicationVndRenamemeItemJSONView) *ApplicationVndRenamemeItemJSONResponseBody {
	if v == nil {
		return nil
	}
	res := &ApplicationVndRenamemeItemJSONResponseBody{
		ID:          v.ID,
		Name:        *v.Name,
		Description: v.Description,
		Icon:        v.Icon,
	}
	if v.Labels != nil {
		res.Labels = make([]string, len(v.Labels))
		for i, val := range v.Labels {
			res.Labels[i] = val
		}
	}

	return res
}

// unmarshalApplicationVndRenamemeItemJSONRequestBodyToItemsApplicationVndRenamemeItemJSON
// builds a value of type *items.ApplicationVndRenamemeItemJSON from a value of
// type *ApplicationVndRenamemeItemJSONRequestBody.
func unmarshalApplicationVndRenamemeItemJSONRequestBodyToItemsApplicationVndRenamemeItemJSON(v *ApplicationVndRenamemeItemJSONRequestBody) *items.ApplicationVndRenamemeItemJSON {
	res := &items.ApplicationVndRenamemeItemJSON{
		ID:          v.ID,
		Name:        *v.Name,
		Description: v.Description,
		Icon:        v.Icon,
	}
	if v.Labels != nil {
		res.Labels = make([]string, len(v.Labels))
		for i, val := range v.Labels {
			res.Labels[i] = val
		}
	}

	return res
}
