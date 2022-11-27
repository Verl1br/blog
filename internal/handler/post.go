package handler

import (
	"net/http"
	"strconv"

	"github.com/dhevve/blog/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type getAllItemsResponse struct {
	Data []model.Post `json:"data"`
}

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var post model.Post

	if err := c.BindJSON(&post); err != nil {
		logrus.Errorf("BindJson Error: %s", err.Error())
		return
	}

	post.UserId = userId

	id, err := h.services.Post.CreatePost(post)
	if err != nil {
		logrus.Errorf("CreatePost Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Errorf("Param Error: %s", err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		logrus.Errorf("GetUserId Error: %s", err.Error())
		return
	}

	post, err := h.services.Post.GetPost(id, userId)
	if err != nil {
		logrus.Errorf("GetPost Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) getPosts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.Errorf("GetUserId Error: %s", err.Error())
		return
	}

	posts, err := h.services.Post.GetPosts(userId)
	if err != nil {
		logrus.Errorf("GetPosts Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: posts,
	})
}

func (h *Handler) deletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err = h.services.Post.DeletePost(id)
	if err != nil {
		logrus.Errorf("DeletePost Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) updatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var post model.UpdatePost

	if err := c.BindJSON(&post); err != nil {
		logrus.Errorf("BindJson Error: %s", err.Error())
		return
	}

	err = h.services.Post.UpdatePost(id, post)
	if err != nil {
		logrus.Errorf("UpdatePost Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
