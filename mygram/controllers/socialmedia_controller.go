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

// SocialMediaInsert godoc
// @Summary Create new Social Media
// @Description Create new Social Media
// @Param data body transports.RequestSocialMedia true "Social Media data"
// @Success 201 {object} transports.ResponseSocialMedia "Social Media data"
// @Failure 400 {object} transports.ResponseSocialMedia
// @Router /socialmedias [post]
// @Tags SocialMedias
func SocialMediaInsert(c *gin.Context) {
	db := lib.DB
	req := transports.RequestSocialMedia{}
	var repoSocialMedia repository.ISocialMedia

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	socialMediaData := models.SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserID:         *lib.GetUserIDFromClaim(),
	}

	if err := repoSocialMedia.InsertSocialMedia(db, &socialMediaData); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	responseData := transports.NewResponseSocialMedia(&socialMediaData)

	transports.SendResponse(c, http.StatusCreated, responseData, nil)
}

// GetListSocialMedias godoc
// @Summary List of Social Media
// @Description List of Social Media
// @Success 200 {object} []transports.RequestSocialMedia List of Social Media
// @Failure 400 {object} transports.ResponseSocialMedia
// @Router /socialmedias [get]
// @Tags SocialMedias
func GetListSocialMedias(c *gin.Context) {
	db := lib.DB
	req := transports.RequestSocialMedia{}
	var repoSocialMedia repository.ISocialMedia

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	socialMediaData, err := repoSocialMedia.GetListSocialMedias(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	transports.SendResponse(c, http.StatusCreated, socialMediaData, nil)
}

// SocialMediaUpdate godoc
// @Summary Update Social Media by id
// @Description Update Social Media by id
// @Param data body transports.RequestSocialMedia true "Social Media data"
// @Success 200 {object} transports.ResponseSocialMedia "Social Media data"
// @Failure 400 {object} transports.ResponseSocialMedia
// @Router /socialmedias/{id} [put]
// @Tags SocialMedias
func SocialMediaUpdate(c *gin.Context) {
	db := lib.DB
	req := transports.RequestSocialMedia{}
	var repoSocialMedia repository.ISocialMedia

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	data, err := repoSocialMedia.GetSocialMediaDataByUserID(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	if data.User.Username != *lib.GetUsernameFromClaim() || data.User.Email != *lib.GetEmailFromClaim() {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	socialMediaData := models.SocialMedia{
		ID:             socialMediaID,
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
	}

	if err := repoSocialMedia.UpdateDataSocialMedia(db, &socialMediaData); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	responseData := transports.NewResponseSocialMedia(&socialMediaData)

	transports.SendResponse(c, http.StatusOK, responseData, nil)
}

// SocialMediaDelete godoc
// @Summary Delete Social Media by id
// @Description Delete Social Media by id
// @Success 200 {object} transports.ResponseSocialMedia
// @Failure 400 {object} transports.ResponseSocialMedia
// @Router /socialmedias/{id} [delete]
// @Tags SocialMedias
func SocialMediaDelete(c *gin.Context) {
	db := lib.DB
	var repoSocialMedia repository.ISocialMedia

	socialMedia, _ := strconv.Atoi(c.Param("socialMediaId"))

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	data, err := repoSocialMedia.GetSocialMediaDataByUserID(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	if data.User.Username != *lib.GetUsernameFromClaim() || data.User.Email != *lib.GetEmailFromClaim() {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	socialMediaData := models.SocialMedia{
		ID: socialMedia,
	}

	if err := repoSocialMedia.DeleteDataSocialMedia(db, &socialMediaData); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	respData := map[string]string{
		"message": "Your social media has been successfully deleted",
	}

	transports.SendCustomResponse(c, http.StatusOK, respData, nil)
}
