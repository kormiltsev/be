package design

import (
	. "goa.design/goa/v3/dsl"
)

var InternalErrorMedia = ResultType("application/vnd.internal.error+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("msgCode")
		Attribute("msg")
		Attribute("attributes")
	})
	Required("msg", "msgCode")
})

var NotFoundErrorMedia = ResultType("application/vnd.not.found.error+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("msgCode")
		Attribute("msg")
		Attribute("attributes")
	})
	Required("msg", "msgCode")
})

var BadRequestErrorMedia = ResultType("application/vnd.bad.request.error+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("msgCode")
		Attribute("msg")
		Attribute("attributes")
	})
	Required("msg", "msgCode")
})
