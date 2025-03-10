// Code generated by goa v3.19.1, DO NOT EDIT.
//
// version HTTP client types
//
// Command:
// $ goa gen github.com/kormiltsev/be/internal/api/design -o ./internal/api

package client

import (
	versionviews "github.com/kormiltsev/be/internal/api/gen/version/views"
)

// VersionResponseBody is the type of the "version" service "version" endpoint
// HTTP response body.
type VersionResponseBody struct {
	// Software version
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// Git commit hash
	Git *string `form:"git,omitempty" json:"git,omitempty" xml:"git,omitempty"`
}

// NewVersionRenamemeVersionOK builds a "version" service "version" endpoint
// result from a HTTP "OK" response.
func NewVersionRenamemeVersionOK(body *VersionResponseBody) *versionviews.RenamemeVersionView {
	v := &versionviews.RenamemeVersionView{
		Version: body.Version,
		Git:     body.Git,
	}

	return v
}
