package migrations

import (
	"github.com/aulianurhady/training/mygram/models"
	"gorm.io/gorm"
)

var ModelMigrations = []interface{}{
	&models.User{},
	&models.SocialMedia{},
	&models.Photo{},
	&models.Comment{},
}

func AutoMigrate(tx *gorm.DB) {
	tx.AutoMigrate(ModelMigrations...)
}
