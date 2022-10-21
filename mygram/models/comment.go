package models

import (
	"time"
)

type Comment struct {
	User  User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Photo Photo `json:"photo,omitempty" gorm:"foreignKey:PhotoID"`

	ID        int        `json:"id,omitempty" gorm:"type:int;primary_key;auto_increment;not_null"`
	UserID    int        `json:"user_id,omitempty" gorm:"type:int"`
	PhotoID   int        `json:"photo_id,omitempty" gorm:"type:int"`
	Message   string     `json:"message,omitempty" gorm:"type:varchar(36);not null"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz"`
}
