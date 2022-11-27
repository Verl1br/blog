package handler

import (
	"fmt"
	"net/http"

	"github.com/dhevve/blog/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		logrus.Fatalf("BindJson Error: %s", err)
		return
	}

	err := h.validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			logrus.Fatalf("Validate Error: %s", err)
		}

		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		logrus.Fatalf("CreateUser Error: %s", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input model.SingInInput

	if err := c.BindJSON(&input); err != nil {
		return
	}

	err := h.validate.Struct(input)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			logrus.Fatalf("Validate Error: %s", err)
		}

		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		logrus.Fatalf("GenerateToken Error: %s", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
