package repository

import (
	"github.com/aulianurhady/training/mygram/models"
	"gorm.io/gorm"
)

type Photo interface {
	InsertPhoto(*gorm.DB, *models.Photo) error
	GetPhotoDataByUserID(*gorm.DB, int) (models.Photo, error)
	GetListPhotos(*gorm.DB, *models.Photo) ([]models.Photo, error)
	UpdateDataPhoto(*gorm.DB, *models.Photo) error
	DeleteDataPhoto(*gorm.DB, *models.Photo) error
}

type IPhoto struct{}

func (IPhoto) InsertPhoto(tx *gorm.DB, data *models.Photo) error {
	return tx.Create(&data).Error
}

func (IPhoto) GetListPhotos(tx *gorm.DB, userID int) ([]models.Photo, error) {
	var data []models.Photo

	if err := tx.Where("user_id = ?", userID).Joins("User").Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (IPhoto) GetPhotoDataByUserID(tx *gorm.DB, userID int) (models.Photo, error) {
	var data models.Photo

	if err := tx.Where("user_id = ?", userID).Joins("User").First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (IPhoto) UpdateDataPhoto(tx *gorm.DB, data *models.Photo) error {
	return tx.Model(&data).Updates(&data).Error
}

func (IPhoto) DeleteDataPhoto(tx *gorm.DB, data *models.Photo) error {
	return tx.Model(&data).Where("id = ?", data.ID).Delete(&data).Error
}
