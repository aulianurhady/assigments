package models

type SocialMedia struct {
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	ID             int    `json:"id,omitempty" gorm:"type:int;primary_key;auto_increment;not_null"`
	Name           string `json:"name,omitempty" gorm:"type:varchar(36);not null"`
	SocialMediaURL string `json:"social_media_url,omitempty" gorm:"type:text;not null"`
	UserID         int    `json:"user_id,omitempty" gorm:"type:int"`
}
