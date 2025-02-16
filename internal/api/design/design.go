package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("renameme", func() {
	Version("1.0")
	Title("RENAMEME")
	Description("RESTful API for handling RENAMEME.")

	HTTP(func() {
		Path("/renameme") // Prefix to all request paths
	})

	Server("renameme", func() {
		Services(
			"items",
			"version",
		)

		Host("localhost", func() {
			URI("https://localhost:8080/renameme")
		})
	})
})
