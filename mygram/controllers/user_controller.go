package controllers

import (
	"net/http"

	"github.com/aulianurhady/training/mygram/lib"
	"github.com/aulianurhady/training/mygram/models"
	"github.com/aulianurhady/training/mygram/repository"
	"github.com/aulianurhady/training/mygram/transports"
	"github.com/gin-gonic/gin"
)

// UserRegister godoc
// @Summary Create new User
// @Description Create new User
// @Param data body transports.RequestUser true "User data"
// @Success 201 {object} transports.RequestUser "User data"
// @Failure 400 {object} transports.Response
// @Router /users/register [post]
// @Tags Users
func UserRegister(c *gin.Context) {
	db := lib.DB
	req := transports.RequestUser{}
	var repoUser repository.IUser

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	hashPasword, err := lib.HashPassword(req.Password)
	if err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	userData := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPasword,
		Age:      req.Age,
	}

	if err := repoUser.Register(db, &userData); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	responseData := transports.NewResponseUser(&userData)

	transports.SendResponse(c, http.StatusCreated, responseData, nil)
}

// UserLogin godoc
// @Summary User Data
// @Description User Data
// @Success 200 {object} transports.Request User Data
// @Failure 400 {object} transports.Response
// @Router /users/login  [post]
// @Tags Users
func UserLogin(c *gin.Context) {
	db := lib.DB
	req := transports.RequestUser{}
	var repoUser repository.IUser

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	userData := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	userData, err := repoUser.Login(db, &userData)
	if err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	if !lib.CheckPasswordHash(req.Password, userData.Password) {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	tokenString, err := lib.BuildJWT(userData.Username, userData.Email)
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	respData := map[string]string{
		"token": tokenString,
	}

	transports.SendCustomResponse(c, http.StatusOK, respData, nil)
}

// UserUpdate godoc
// @Summary Update User by email and username
// @Description Update User by email and username
// @Param data body transports.Request true "User data"
// @Success 200 {object} transports.Request "User data"
// @Failure 400 {object} transports.Response
// @Router /users [put]
// @Tags Users
func UserUpdate(c *gin.Context) {
	db := lib.DB
	req := transports.RequestUser{}
	var repoUser repository.IUser

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	if err := lib.Auth(c.GetHeader("Authorization"), req.Username, req.Email); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	userData := models.User{
		Username: lib.GetUsernameFromClaim(),
		Email:    lib.GetEmailFromClaim(),
	}

	if err := repoUser.UpdateDataUser(db, &userData); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	responseData := transports.NewResponseUser(&userData)

	transports.SendResponse(c, http.StatusOK, responseData, nil)
}

// UserDelete godoc
// @Summary Delete User by username and email
// @Description Delete User by username and email
// @Param data body transports.Request true "User data"
// @Success 200 {object} transports.Response
// @Failure 400 {object} transports.Response
// @Router /users [delete]
// @Tags Users
func UserDelete(c *gin.Context) {
	db := lib.DB
	req := transports.RequestUser{}
	var repoUser repository.IUser

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	if err := lib.Auth(c.GetHeader("Authorization"), req.Username, req.Email); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	userData := models.User{
		Username: lib.GetUsernameFromClaim(),
		Email:    lib.GetEmailFromClaim(),
	}

	if err := repoUser.DeleteDataUser(db, &userData); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	respData := map[string]string{
		"message": "Your account has been successfully deleted",
	}

	transports.SendCustomResponse(c, http.StatusOK, respData, nil)
}
