package repository

import (
	"github.com/aulianurhady/training/mygram/models"
	"gorm.io/gorm"
)

type Comment interface {
	InsertComment(*gorm.DB, *models.Comment) error
	GetCommentDataByUserID(*gorm.DB, int) (models.Comment, error)
	GetListComments(*gorm.DB, *models.Comment) ([]models.Comment, error)
	UpdateDataComment(*gorm.DB, *models.Comment) error
	DeleteDataComment(*gorm.DB, *models.Comment) error
}

type IComment struct{}

func (IComment) InsertComment(tx *gorm.DB, data *models.Comment) error {
	return tx.Create(&data).Error
}

func (IComment) GetListComments(tx *gorm.DB, userID int) ([]models.Comment, error) {
	var data []models.Comment

	if err := tx.Where("user_id = ?", userID).Joins("User").Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (IComment) GetCommentDataByUserID(tx *gorm.DB, userID int) (models.Comment, error) {
	var data models.Comment

	if err := tx.Where("user_id = ?", userID).Joins("User").First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (IComment) UpdateDataComment(tx *gorm.DB, data *models.Comment) error {
	return tx.Model(&data).Updates(&data).Error
}

func (IComment) DeleteDataComment(tx *gorm.DB, data *models.Comment) error {
	return tx.Model(&data).Where("id = ?", data.ID).Delete(&data).Error
}
