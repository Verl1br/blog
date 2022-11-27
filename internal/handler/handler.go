package handler

import (
	"github.com/dhevve/blog/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	services *service.Service
	validate *validator.Validate
}

func NewHandler(service *service.Service, validate *validator.Validate) *Handler {
	return &Handler{services: service, validate: validate}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/sing-in", h.signIn)
	}

	post := router.Group("/post", h.userIdentity)
	{
		post.POST("/create-post", h.createPost)
		post.GET("/get-post/:id", h.getPost)
		post.GET("/get-posts", h.getPosts)
		post.DELETE("/delete-post/:id", h.deletePost)
		post.PUT("/update-post/:id", h.updatePost)
	}

	return router
}
