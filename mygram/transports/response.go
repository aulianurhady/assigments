package transports

import (
	"time"

	"github.com/aulianurhady/training/mygram/models"
	"github.com/gin-gonic/gin"
)

type responseUser struct {
	Age       int        `json:"age,omitempty"`
	Email     string     `json:"email,omitempty"`
	ID        int        `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type responsePhoto struct {
	ID        int        `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Caption   string     `json:"caption,omitempty"`
	PhotoURL  string     `json:"photo_url,omitempty"`
	UserID    int        `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func NewResponseUser(data *models.User) *responseUser {
	return &responseUser{
		Age:       data.Age,
		Email:     data.Email,
		ID:        data.ID,
		Username:  data.Username,
		UpdatedAt: data.UpdatedAt,
	}
}

func NewResponsePhoto(data *models.Photo) *responsePhoto {
	return &responsePhoto{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoURL:  data.PhotoURL,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}
}

func SendResponse(c *gin.Context, statusCode int, data interface{}, err error) {
	c.Header("Content-Type", "application/json")

	c.JSON(statusCode, gin.H{
		"data": data,
	})
}

func SendCustomResponse(c *gin.Context, statusCode int, mapData interface{}, err error) {
	c.Header("Content-Type", "application/json")

	c.JSON(statusCode, mapData)
}
