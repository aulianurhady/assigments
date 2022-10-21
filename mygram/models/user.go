package models

import (
	"time"
)

type User struct {
	ID        int        `json:"id,omitempty" gorm:"type:int;primary_key;auto_increment;not_null"`
	Username  string     `json:"username,omitempty" gorm:"type:varchar(36);not null;unique"`
	Email     string     `json:"email,omitempty" gorm:"type:varchar(36);not null;unique"`
	Password  string     `json:"password,omitempty" gorm:"type:varchar(36);not null" validate:"required,min=6"`
	Age       int        `json:"order_id,omitempty" gorm:"type:int;not null" validate:"required,numeric,min=9"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz"`
}
