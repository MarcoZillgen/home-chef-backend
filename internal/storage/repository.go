package storage

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateItem(item StorageItem) error {
	return r.db.Create(&item).Error
}

func (r *Repository) GetItems() ([]StorageItem, error) {
	var items []StorageItem
	err := r.db.Find(&items).Error
	return items, err
}

func (r *Repository) GetItemByID(id string) (StorageItem, error) {
	var item StorageItem
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *Repository) UpdateItem(item StorageItem) error {
	result := r.db.Save(&item)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no item found with id %s", item.ID)
	}
	return nil
}

func (r *Repository) DeleteItem(id string) error {
	result := r.db.Delete(&StorageItem{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no item found with id %s", id)
	}
	return nil
}
