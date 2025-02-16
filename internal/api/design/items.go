package design

import (
	. "goa.design/goa/v3/dsl"
)

var ItemMedia = Type("application/vnd.renameme.item+json", func() {
	Attribute("id", String, "Item identifier")
	Attribute("name", String, "Item name", func() {
		MinLength(1)
		MaxLength(255)
	})
	Attribute("description", String, "Item description")
	Attribute("labels", ArrayOf(String), "Item labels")
	Attribute("icon", String, "Application icon in base64 format")
	Required("name")
})

var ItemsMedia = ResultType("application/vnd.renameme.items+json", func() {
	Description("Items")
	ContentType("application/json")
	Attributes(func() {
		Attribute("items", ArrayOf(ItemMedia), "Items collection")
	})
})

var _ = Service("items", func() {
	Error("not_found", NotFoundErrorMedia, "is a common error response for not found")
	Error("bad_request", BadRequestErrorMedia, "is a common error response for bad request")
	Error("internal", InternalErrorMedia, "is a common error response for internal error")

	Method("list", func() {
		Description(`Return list of all items`)

		Result(ItemsMedia)

		HTTP(func() {
			GET("/items")

			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
			Response("internal", StatusInternalServerError)

			Meta("swagger:summary", "Return list of items")
		})
	})

	Method("create", func() {
		Payload(func() {
			Field(1, "item", ItemMedia)
			Required("item")
		})
		HTTP(func() {
			POST("/items")

			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
			Response("internal", StatusInternalServerError)

			Meta("swagger:summary", "Create new item")
		})
	})

	Method("delete", func() {

		Payload(func() {
			Field(1, "id", String, "Item ID", func() {
				Example("asdf-asdf-asdf")
			})
			Required("id")
		})

		HTTP(func() {
			DELETE("/items/{id}")

			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
			Response("bad_request", StatusBadRequest)
			Response("internal", StatusInternalServerError)

			Meta("swagger:summary", "Delete item")
		})
	})
})
