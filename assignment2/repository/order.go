package repository

import (
	"github.com/aulianurhady/training/assignment2/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateDataOrder(*gorm.DB, *models.Orders) error
	GetListOrder(*gorm.DB) ([]models.Orders, error)
	UpdateDataOrder(*gorm.DB, *models.Orders) error
	GetDataOrderByID(*gorm.DB, int) (models.Orders, error)
	DeleteDataOrder(*gorm.DB, *models.Orders) error
}

func CreateDataOrder(tx *gorm.DB, data *models.Orders) error {
	return tx.Create(&data).Error
}

func GetListOrder(tx *gorm.DB) ([]models.Orders, error) {
	var data []models.Orders

	if err := tx.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func GetDataOrderByID(tx *gorm.DB, id int) (models.Orders, error) {
	var data models.Orders

	if err := tx.Where("order_id = ?", id).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func UpdateDataOrder(tx *gorm.DB, data *models.Orders) error {
	return tx.Model(&data).Updates(&data).Error
}

func DeleteDataOrder(tx *gorm.DB, data *models.Orders) error {
	return tx.Delete(&data).Error
}
