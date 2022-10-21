package controllers

import (
	"net/http"
	"strconv"

	"github.com/aulianurhady/training/mygram/lib"
	"github.com/aulianurhady/training/mygram/models"
	"github.com/aulianurhady/training/mygram/repository"
	"github.com/aulianurhady/training/mygram/transports"
	"github.com/gin-gonic/gin"
)

// PhotoInsert godoc
// @Summary Create new Photos
// @Description Create new Photos
// @Param data body transports.RequestUser true "User data"
// @Success 201 {object} transports.RequestUser "User data"
// @Failure 400 {object} transports.Response
// @Router /users/register [post]
// @Tags Users
func PhotoInsert(c *gin.Context) {
	db := lib.DB
	req := transports.RequestPhoto{}
	var repoPhoto repository.IPhoto

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	photoData := models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
		UserID:   *lib.GetUserIDFromClaim(),
	}

	if err := repoPhoto.InsertPhoto(db, &photoData); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	responseData := transports.NewResponsePhoto(&photoData)

	transports.SendResponse(c, http.StatusCreated, responseData, nil)
}

// GetListPhotos godoc
// @Summary List of Photos
// @Description List of Photos
// @Success 200 {object} []transports.Request List of Photos
// @Failure 400 {object} transports.Response
// @Router /photos [get]
// @Tags Photos
func GetListPhotos(c *gin.Context) {
	db := lib.DB
	req := transports.RequestPhoto{}
	var repoPhoto repository.IPhoto

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	photoData, err := repoPhoto.GetListPhotos(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	transports.SendResponse(c, http.StatusCreated, photoData, nil)
}

// PhotoUpdate godoc
// @Summary Update Photo by id
// @Description Update Photo by id
// @Param data body transports.Request true "Photo data"
// @Success 200 {object} transports.Request "Photo data"
// @Failure 400 {object} transports.Response
// @Router /photos/{id} [put]
// @Tags Users
func PhotoUpdate(c *gin.Context) {
	db := lib.DB
	req := transports.RequestPhoto{}
	var repoPhoto repository.IPhoto

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	photoID, _ := strconv.Atoi(c.Param("photoID"))

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	data, err := repoPhoto.GetPhotoDataByUserID(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	if data.User.Username != *lib.GetUsernameFromClaim() || data.User.Email != *lib.GetEmailFromClaim() {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	photoData := models.Photo{
		ID:       photoID,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
	}

	if err := repoPhoto.UpdateDataPhoto(db, &photoData); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	responseData := transports.NewResponsePhoto(&photoData)

	transports.SendResponse(c, http.StatusOK, responseData, nil)
}

// PhotoDelete godoc
// @Summary Delete Photo by id
// @Description Delete Photo by id
// @Param data body transports.Request true "Photo data"
// @Success 200 {object} transports.Response
// @Failure 400 {object} transports.Response
// @Router /photos/{id} [delete]
// @Tags Photos
func PhotoDelete(c *gin.Context) {
	db := lib.DB
	var repoPhoto repository.IPhoto

	photoID, _ := strconv.Atoi(c.Param("photoID"))

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	data, err := repoPhoto.GetPhotoDataByUserID(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	if data.User.Username != *lib.GetUsernameFromClaim() || data.User.Email != *lib.GetEmailFromClaim() {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	photoData := models.Photo{
		ID: photoID,
	}

	if err := repoPhoto.DeleteDataPhoto(db, &photoData); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	respData := map[string]string{
		"message": "Your photo has been successfully deleted",
	}

	transports.SendCustomResponse(c, http.StatusOK, respData, nil)
}
