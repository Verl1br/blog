package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) getFriends(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Errorf("Param Error: %s", err.Error())
		return
	}
	logrus.Info("GetFriends!")
	friends := h.services.Friend.GetFriends(id)

	c.JSON(http.StatusOK, friends)
}

func (h *Handler) createFriends(c *gin.Context) {
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

	err = h.services.Friend.CreateFriends(userId, id)
	if err != nil {
		logrus.Fatalf("CreateFriends Error: %s", err.Error())
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) deleteFriends(c *gin.Context) {
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

	err = h.services.Friend.DeleteFriend(userId, id)
	if err != nil {
		logrus.Fatalf("DeleteFriend Error: %s", err.Error())
	}

	c.JSON(http.StatusOK, "ok")
}
