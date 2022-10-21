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

func NewResponseUser(data *models.User) *responseUser {
	return &responseUser{
		Age:       data.Age,
		Email:     data.Email,
		ID:        data.ID,
		Username:  data.Username,
		UpdatedAt: data.UpdatedAt,
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
