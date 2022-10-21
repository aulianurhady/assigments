package transports

import (
	"time"

	"github.com/aulianurhady/training/mygram/models"
	"github.com/gin-gonic/gin"
)

type ResponseUser struct {
	Age       int        `json:"age,omitempty"`
	Email     string     `json:"email,omitempty"`
	ID        int        `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ResponsePhoto struct {
	ID        int        `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Caption   string     `json:"caption,omitempty"`
	PhotoURL  string     `json:"photo_url,omitempty"`
	UserID    int        `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type ResponseComment struct {
	ID        int        `json:"id,omitempty"`
	UserID    int        `json:"user_id,omitempty"`
	PhotoID   int        `json:"photo_id,omitempty"`
	Message   string     `json:"message,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func NewResponseUser(data *models.User) *ResponseUser {
	return &ResponseUser{
		Age:       data.Age,
		Email:     data.Email,
		ID:        data.ID,
		Username:  data.Username,
		UpdatedAt: data.UpdatedAt,
	}
}

func NewResponsePhoto(data *models.Photo) *ResponsePhoto {
	return &ResponsePhoto{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoURL:  data.PhotoURL,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
	}
}

func NewResponseComment(data *models.Comment) *ResponseComment {
	return &ResponseComment{
		ID:        data.ID,
		PhotoID:   data.PhotoID,
		Message:   data.Message,
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
