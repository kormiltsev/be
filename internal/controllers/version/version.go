package version

import (
	"context"

	goasvc "github.com/kormiltsev/be/internal/api/gen/version"
	"github.com/kormiltsev/be/version"
)

// Controller implements the version resource.
type Controller struct {
}

// NewController creates a version controller.
func NewController() *Controller {
	return &Controller{}
}

// Version runs the version action.
func (c *Controller) Version(ctx context.Context) (*goasvc.RenamemeVersion, error) {
	return &goasvc.RenamemeVersion{
		Version: version.Version,
		Git:     &version.GitCommit,
	}, nil
}
