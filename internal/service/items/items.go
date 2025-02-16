package items

import (
	"context"
	"fmt"
	"log/slog"

	tomb "gopkg.in/tomb.v2"

	"github.com/kormiltsev/be/internal/model"
)

type Dal interface {
	CreateItem(ctx context.Context, item *model.Item) error
	GetItemByID(ctx context.Context, id string) (*model.Item, error)
	GetItems(ctx context.Context) ([]*model.Item, error)
	UpdateItem(ctx context.Context, item *model.Item) error
	DeleteItem(ctx context.Context, id string) error
}

type Service struct {
	t   *tomb.Tomb
	log *slog.Logger
	dal Dal
}

func New(t *tomb.Tomb, log *slog.Logger, dal Dal) *Service {
	return &Service{t, log.With(slog.String("layer", "service")), dal}
}

func (s *Service) List(ctx context.Context) ([]*model.Item, error) {
	items, err := s.dal.GetItems(ctx)
	if err != nil {
		s.log.Error("List of Items error", slog.String("error", err.Error()))
		return []*model.Item{}, fmt.Errorf("get list returns: %w", err)
	}
	return items, nil
	// return []*model.Item{{ID: "id_0001-0001", Name: "NameItem"}, {ID: "id_0001-0002", Name: "NameItem2"}}, nil
}

func (s *Service) Create(ctx context.Context, item model.Item) error {
	return s.dal.CreateItem(ctx, &item)
	// return &model.Item{ID: "id_0001-0001", Name: "NameItem"}, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.dal.DeleteItem(ctx, id)
}
