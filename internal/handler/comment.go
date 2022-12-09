package handler

import (
	"net/http"
	"strconv"

	"github.com/dhevve/blog/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Errorf("Param Error: %s", err.Error())
		return
	}

	var comment model.Comment

	if err := c.BindJSON(&comment); err != nil {
		logrus.Errorf("BindJson Error: %s", err.Error())
		return
	}

	err = h.validate.Struct(comment)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logrus.Info(err.Error())
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			logrus.Fatalf("Validate Error: %s", err)
		}

		return
	}

	comment.UserId = userId
	comment.PostId = id

	commentId, err := h.services.Сomment.CreateComment(comment)
	if err != nil {
		logrus.Errorf("CreateComment Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, commentId)
}

func (h *Handler) getComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Errorf("Param Error: %s", err.Error())
		return
	}

	post, err := h.services.Сomment.GetComment(id)
	if err != nil {
		logrus.Errorf("GetComment Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) getComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	comments, err := h.services.Сomment.GetComments(id)
	if err != nil {
		logrus.Errorf("GetComments Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *Handler) deleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err = h.services.Сomment.DeleteComment(id)
	if err != nil {
		logrus.Errorf("DeleteСomment Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) updateComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var comment model.UpdateComment

	if err := c.BindJSON(&comment); err != nil {
		logrus.Errorf("BindJson Error: %s", err.Error())
		return
	}

	err = h.validate.Struct(comment)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logrus.Info(err.Error())
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			logrus.Fatalf("Validate Error: %s", err)
		}

		return
	}

	err = h.services.Сomment.UpdateComment(id, comment)
	if err != nil {
		logrus.Errorf("UpdateСomment Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
