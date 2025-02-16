package design

import (
	. "goa.design/goa/v3/dsl"
)

var VersionType = Type("version", func() {
	Attribute("version", String, "Software version", func() {
		Example("1.2.3")
	})
	Attribute("git", String, "Git commit hash", func() {
		Example("cd243db4")
	})
	Required("version")
})

var VersionMedia = ResultType("application/vnd.renameme.version+json", func() {
	Description("Version")
	ContentType("application/json")
	Attributes(func() {
		Attribute("version", String, "Software version", func() {
			Example("1.2.3")
		})
		Attribute("git", String, "Git commit hash", func() {
			Example("cd243db4")
		})
		Required("version")
	})
})

var _ = Service("version", func() {
	Description("Get Version")
	Method("version", func() {
		Result(VersionMedia)

		HTTP(func() {
			GET("/version")
			Response(StatusOK)
			Meta("swagger:summary", "Get renameme version")
		})
	})
})
