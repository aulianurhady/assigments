package models

type Items struct {
	ItemID      int    `json:"item_id,omitempty" gorm:"type:int;primary_key;auto_increment;not_null"`
	ItemCode    string `json:"item_code,omitempty" gorm:"type:varchar(36)"`
	Description string `json:"description,omitempty" gorm:"type:varchar(36)"`
	Quantity    int    `json:"quantity,omitempty" gorm:"type:int"`
	OrderID     int    `json:"order_id,omitempty" gorm:"type:int"`
}
