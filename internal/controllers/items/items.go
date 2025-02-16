package items

import (
	"context"
	"log/slog"
	"strconv"

	goasvc "github.com/kormiltsev/be/internal/api/gen/items"
	"github.com/kormiltsev/be/internal/model"
)

type ItemService interface {
	List(context.Context) ([]*model.Item, error)
	Create(context.Context, model.Item) error
	Delete(context.Context, string) error
}

// Controller implements the version resource.
type Controller struct {
	itemService ItemService

	log *slog.Logger
}

// NewController creates a version controller.
func NewController(service ItemService, log *slog.Logger) *Controller {
	return &Controller{
		itemService: service,
		log:         log.With(slog.String("layer", "controller")),
	}
}

// List items.
func (c *Controller) List(ctx context.Context) (*goasvc.RenamemeItems, error) {
	items, err := c.itemService.List(ctx)
	if err != nil {
		return nil, err
	}
	c.log.Debug("list of items", slog.String("quantity", strconv.Itoa(len(items))))
	return itemsToResponse(items)
}

// Create item.
func (c *Controller) Create(ctx context.Context, p *goasvc.CreatePayload) error {
	err := c.itemService.Create(ctx, model.Item{
		Name: p.Item.Name,
	})
	if err != nil {
		c.log.Error("create", "error:", err)
		return err
	}
	return nil
}

// Delete item.
func (c *Controller) Delete(ctx context.Context, p *goasvc.DeletePayload) error {
	err := c.itemService.Delete(ctx, p.ID)
	if err != nil {
		c.log.Error("delete", "error:", err)
		return err
	}
	return nil
}
