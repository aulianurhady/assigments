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

// CommentInsert godoc
// @Summary Create new Comment
// @Description Create new Comment
// @Param data body transports.RequestComment true "Comment data"
// @Success 201 {object} transports.ResponseComment "Comment data"
// @Failure 400 {object} transports.ResponseComment
// @Router /comments [post]
// @Tags Comments
func CommentInsert(c *gin.Context) {
	db := lib.DB
	req := transports.RequestComment{}
	var repoComment repository.IComment

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	commentData := models.Comment{
		PhotoID: req.PhotoID,
		Message: req.Message,
		UserID:  *lib.GetUserIDFromClaim(),
	}

	if err := repoComment.InsertComment(db, &commentData); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	responseData := transports.NewResponseComment(&commentData)

	transports.SendResponse(c, http.StatusCreated, responseData, nil)
}

// GetListComments godoc
// @Summary List of Comments
// @Description List of Comments
// @Success 200 {object} []transports.RequestComment List of Comments
// @Failure 400 {object} transports.ResponseComment
// @Router /comments [get]
// @Tags Comments
func GetListComments(c *gin.Context) {
	db := lib.DB
	req := transports.RequestPhoto{}
	var repoComment repository.IPhoto

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	photoData, err := repoComment.GetListPhotos(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	transports.SendResponse(c, http.StatusCreated, photoData, nil)
}

// CommentUpdate godoc
// @Summary Update Comment by id
// @Description Update Comment by id
// @Param data body transports.RequestComment true "Comment data"
// @Success 200 {object} transports.ResponseComment "Comment data"
// @Failure 400 {object} transports.ResponseComment
// @Router /comments/{id} [put]
// @Tags Comments
func CommentUpdate(c *gin.Context) {
	db := lib.DB
	req := transports.RequestComment{}
	var repoComment repository.IComment

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	commentID, _ := strconv.Atoi(c.Param("commentId"))

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	data, err := repoComment.GetCommentDataByUserID(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	if data.User.Username != *lib.GetUsernameFromClaim() || data.User.Email != *lib.GetEmailFromClaim() {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	commentData := models.Comment{
		ID:      commentID,
		Message: req.Message,
	}

	if err := repoComment.UpdateDataComment(db, &commentData); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	responseData := transports.NewResponseComment(&commentData)

	transports.SendResponse(c, http.StatusOK, responseData, nil)
}

// CommentDelete godoc
// @Summary Delete Comment by id
// @Description Delete Comment by id
// @Param data body transports.RequestComment true "Comment data"
// @Success 200 {object} transports.ResponseComment
// @Failure 400 {object} transports.ResponseComment
// @Router /comments/{id} [delete]
// @Tags Comments
func CommentDelete(c *gin.Context) {
	db := lib.DB
	var repoComment repository.IComment

	commentID, _ := strconv.Atoi(c.Param("commentId"))

	if err := lib.Auth(c.GetHeader("Authorization")); err != nil {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	data, err := repoComment.GetCommentDataByUserID(db, *lib.GetUserIDFromClaim())
	if err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	if data.User.Username != *lib.GetUsernameFromClaim() || data.User.Email != *lib.GetEmailFromClaim() {
		transports.SendResponse(c, http.StatusBadRequest, nil, err)
		return
	}

	commentData := models.Comment{
		ID: commentID,
	}

	if err := repoComment.DeleteDataComment(db, &commentData); err != nil {
		transports.SendResponse(c, http.StatusInternalServerError, nil, err)
		return
	}

	respData := map[string]string{
		"message": "Your comment has been successfully deleted",
	}

	transports.SendCustomResponse(c, http.StatusOK, respData, nil)
}
