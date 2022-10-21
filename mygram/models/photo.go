package models

import (
	"time"
)

type Photo struct {
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	ID        int        `json:"id,omitempty" gorm:"type:int;primary_key;auto_increment;not_null"`
	Title     string     `json:"title,omitempty" gorm:"type:varchar(36);not null"`
	Caption   string     `json:"caption,omitempty" gorm:"type:varchar(36)"`
	PhotoURL  string     `json:"photo_url,omitempty" gorm:"type:varchar(36);not null"`
	UserID    int        `json:"user_id,omitempty" gorm:"type:int"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz"`
}
