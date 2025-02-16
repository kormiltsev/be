package dal

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/kormiltsev/be/internal/errs"
	"github.com/kormiltsev/be/internal/model"

	"gorm.io/gorm"
)

type ItemDAL struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      // Set to current time if it is zero on creating
	Updated   int64          `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	Created   int64          `gorm:"autoCreateTime"`      // Use unix seconds as creating time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name string `gorm:"size:255;not null"`
	// Metadata   datatypes.JSON    `gorm:"type:json"`       // Store JSON in SQLite // metadataJSON, err := json.Marshal(metadata)
	// Parameters map[string]string `gorm:"serializer:json"` // Automatically converts map to JSON

	OriginalData *model.Item `gorm:"serializer:json"`
}

func init() {
	tables = append(tables, &ItemDAL{})
}

func (dal *Dal) CreateItem(ctx context.Context, item *model.Item) error {
	itemdal := modelToDalNoID(item)
	if err := dal.db.WithContext(ctx).Create(&itemdal).Error; err != nil {
		return err
	}
	return nil
}

// GetItemByID fetches a user by ID
func (dal *Dal) GetItemByID(ctx context.Context, id string) (*model.Item, error) {
	uintid := convertID(id)
	if uintid == 0 {
		return nil, fmt.Errorf("%w id:%s", errs.ErrWrongIdFormat, id)
	}

	var itemdal ItemDAL
	if err := dal.db.WithContext(ctx).First(&itemdal, uintid).Error; err != nil {
		return nil, err
	}
	return dalToModel(&itemdal), nil
}

// GetItems retrieves all users
func (dal *Dal) GetItems(ctx context.Context) ([]*model.Item, error) {
	var itemsdal []*ItemDAL
	if err := dal.db.WithContext(ctx).Find(&itemsdal).Error; err != nil {
		return nil, err
	}
	return dalsToModels(itemsdal), nil
}

// TODO use tx
// UpdateItem updates a item's details
func (dal *Dal) UpdateItem(ctx context.Context, item *model.Item) error {
	uintid := convertID(item.ID)
	if uintid == 0 {
		return fmt.Errorf("%w id:%s", errs.ErrWrongIdFormat, item.ID)
	}

	itemdal, err := dal.GetItemByID(ctx, item.ID)
	if err != nil {
		return err
	}
	itemdal.Name = item.Name
	return dal.db.WithContext(ctx).Save(itemdal).Error
}

// DeleteItem deletes a item by ID
func (dal *Dal) DeleteItem(ctx context.Context, id string) error {
	uintid := convertID(id)
	if uintid == 0 {
		return fmt.Errorf("%w id:%s", errs.ErrWrongIdFormat, id)
	}

	return dal.db.WithContext(ctx).Delete(&ItemDAL{}, id).Error
}

func dalsToModels(dals []*ItemDAL) []*model.Item {
	res := make([]*model.Item, len(dals))
	for i := range dals {
		res[i] = dalToModel(dals[i])
	}
	return res
}

func dalToModel(dal *ItemDAL) *model.Item {
	res := model.Item{
		ID:   strconv.FormatUint(uint64(dal.ID), 10),
		Name: dal.Name,
	}
	return &res
}

func modelToDalNoID(item *model.Item) *ItemDAL {
	res := ItemDAL{
		Name: item.Name,
	}
	return &res
}

func convertID(id string) uint {
	intid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0
	}
	return uint(intid)
}
