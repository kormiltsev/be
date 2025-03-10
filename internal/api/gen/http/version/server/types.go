// Code generated by goa v3.19.1, DO NOT EDIT.
//
// version HTTP server types
//
// Command:
// $ goa gen github.com/kormiltsev/be/internal/api/design -o ./internal/api

package server

import (
	versionviews "github.com/kormiltsev/be/internal/api/gen/version/views"
)

// VersionResponseBody is the type of the "version" service "version" endpoint
// HTTP response body.
type VersionResponseBody struct {
	// Software version
	Version string `form:"version" json:"version" xml:"version"`
	// Git commit hash
	Git *string `form:"git,omitempty" json:"git,omitempty" xml:"git,omitempty"`
}

// NewVersionResponseBody builds the HTTP response body from the result of the
// "version" endpoint of the "version" service.
func NewVersionResponseBody(res *versionviews.RenamemeVersionView) *VersionResponseBody {
	body := &VersionResponseBody{
		Version: *res.Version,
		Git:     res.Git,
	}
	return body
}
