package api

import (
	_ "goa.design/goa/v3/codegen"           //nolint:revive
	_ "goa.design/goa/v3/codegen/generator" //nolint:revive
)

//go:generate goa gen github.com/kormiltsev/be/internal/api/design -o ./internal/api
