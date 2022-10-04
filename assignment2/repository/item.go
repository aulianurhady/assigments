package repository

import (
	"github.com/aulianurhady/training/assignment2/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateDataItem(*gorm.DB, *models.Items) error
	GetListItemByID(*gorm.DB, int) ([]models.Items, error)
	GetDataItemByID(*gorm.DB, int) ([]models.Items, error)
	UpdateDataItem(*gorm.DB, *models.Items) error
	DeleteDataItem(*gorm.DB, *models.Items) error
}

func CreateDataItem(tx *gorm.DB, data *models.Items) error {
	return tx.Create(&data).Error
}

func GetListItemByID(tx *gorm.DB, id int) ([]models.Items, error) {
	var data []models.Items

	if err := tx.Where("order_id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func GetDataItemByID(tx *gorm.DB, id int) (models.Items, error) {
	var data models.Items

	if err := tx.Where("item_id = ?", id).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func UpdateDataItem(tx *gorm.DB, data *models.Items) error {
	return tx.Model(&data).Updates(&data).Error
}

func DeleteDataItem(tx *gorm.DB, data *models.Items) error {
	return tx.Delete(&data).Error
}
