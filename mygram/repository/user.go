package repository

import (
	"github.com/aulianurhady/training/mygram/models"
	"gorm.io/gorm"
)

type User interface {
	Register(*gorm.DB, *models.User) error
	Login(*gorm.DB, *models.User) (models.User, error)
	UpdateDataUser(*gorm.DB, *models.User) error
	DeleteDataUser(*gorm.DB, *models.User) error
}

type IUser struct{}

func (IUser) Register(tx *gorm.DB, data *models.User) error {
	return tx.Create(&data).Error
}

func (IUser) Login(tx *gorm.DB, dataReq *models.User) (models.User, error) {
	var data models.User

	if err := tx.Where("email = ?", dataReq.Email).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (IUser) UpdateDataUser(tx *gorm.DB, data *models.User) error {
	return tx.Model(&data).Updates(&data).Error
}

func (IUser) DeleteDataUser(tx *gorm.DB, data *models.User) error {
	return tx.Model(&data).Delete(&data).Error
}
