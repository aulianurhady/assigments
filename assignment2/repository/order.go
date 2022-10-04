package repository

import (
	"github.com/aulianurhady/training/assignment2/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateDataOrder(*gorm.DB, *models.Orders) error
	GetDataOrder(*gorm.DB) ([]models.Orders, error)
}

func CreateDataOrder(tx *gorm.DB, data *models.Orders) error {
	return tx.Create(&data).Error
}

func GetDataOrder(tx *gorm.DB) ([]models.Orders, error) {
	var data []models.Orders

	if err := tx.Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
