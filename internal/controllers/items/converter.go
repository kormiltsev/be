package items

import (
	goasvc "github.com/kormiltsev/be/internal/api/gen/items"
	"github.com/kormiltsev/be/internal/model"
	"github.com/kormiltsev/be/internal/utils"
)

func itemsToResponse(items []*model.Item) (*goasvc.RenamemeItems, error) {
	resp := []*goasvc.ApplicationVndRenamemeItemJSON{}
	for _, item := range items {
		resp = append(resp, &goasvc.ApplicationVndRenamemeItemJSON{
			ID:   utils.String(item.ID),
			Name: item.Name,
		})
	}
	return &goasvc.RenamemeItems{Items: resp}, nil
}
