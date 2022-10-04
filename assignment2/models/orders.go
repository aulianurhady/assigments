package models

import (
	"time"
)

type Orders struct {
	OrderID      int       `json:"order_id,omitempty" gorm:"type:int;primary_key;auto_increment;not_null"`
	CustomerName string    `json:"customer_name,omitempty" gorm:"type:varchar(36)"`
	OrderedAt    time.Time `json:"ordered_at,omitempty" gorm:"type:timestamptz"`
}
