package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) getNews(c *gin.Context) {
	/*userId, err := getUserId(c)
	if err != nil {
		logrus.Errorf("GetUserId Error: %s", err.Error())
		return
	}*/

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Errorf("Param Error: %s", err.Error())
		return
	}

	posts, err := h.services.NewsFeed.GetNews(id, h.ctx)
	if err != nil {
		logrus.Errorf("GetNews Error: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}
