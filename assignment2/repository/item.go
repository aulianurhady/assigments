package repository

import (
	"github.com/aulianurhady/training/assignment2/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateDataItem(*gorm.DB, *models.Items) error
	GetDataItemByID(*gorm.DB, int) ([]models.Items, error)
}

func CreateDataItem(tx *gorm.DB, data *models.Items) error {
	return tx.Create(&data).Error
}

func GetDataItemByID(tx *gorm.DB, id int) ([]models.Items, error) {
	var data []models.Items

	if err := tx.Where("order_id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
