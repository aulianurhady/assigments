package repository

import (
	"github.com/aulianurhady/training/mygram/models"
	"gorm.io/gorm"
)

type SocialMedia interface {
	InsertSocialMedia(*gorm.DB, *models.SocialMedia) error
	GetSocialMediaDataByUserID(*gorm.DB, int) (models.SocialMedia, error)
	GetListSocialMedias(*gorm.DB, *models.SocialMedia) ([]models.SocialMedia, error)
	UpdateDataSocialMedia(*gorm.DB, *models.SocialMedia) error
	DeleteDataSocialMedia(*gorm.DB, *models.SocialMedia) error
}

type ISocialMedia struct{}

func (ISocialMedia) InsertSocialMedia(tx *gorm.DB, data *models.SocialMedia) error {
	return tx.Create(&data).Error
}

func (ISocialMedia) GetListSocialMedias(tx *gorm.DB, userID int) ([]models.SocialMedia, error) {
	var data []models.SocialMedia

	if err := tx.Where("user_id = ?", userID).Joins("User").Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (ISocialMedia) GetSocialMediaDataByUserID(tx *gorm.DB, userID int) (models.SocialMedia, error) {
	var data models.SocialMedia

	if err := tx.Where("user_id = ?", userID).Joins("User").First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (ISocialMedia) UpdateDataSocialMedia(tx *gorm.DB, data *models.SocialMedia) error {
	return tx.Model(&data).Updates(&data).Error
}

func (ISocialMedia) DeleteDataSocialMedia(tx *gorm.DB, data *models.SocialMedia) error {
	return tx.Model(&data).Where("id = ?", data.ID).Delete(&data).Error
}
