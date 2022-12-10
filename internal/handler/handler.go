package handler

import (
	"context"

	"github.com/dhevve/blog/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	services *service.Service
	validate *validator.Validate
	ctx      context.Context
}

func NewHandler(service *service.Service, validate *validator.Validate, ctx context.Context) *Handler {
	return &Handler{
		services: service,
		validate: validate,
		ctx:      ctx,
	}
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

	comment := router.Group("/comment", h.userIdentity)
	{
		comment.POST("/create-comment/:id", h.createComment)
		comment.GET("/get-comment/:id", h.getComment)
		comment.GET("/get-comments/:id", h.getComments)
		comment.DELETE("/delete-comment/:id", h.deleteComment)
		comment.PUT("/update-comment/:id", h.updateComment)
	}

	friend := router.Group("/friend", h.userIdentity)
	{
		friend.GET("/get-friends/:id", h.getFriends)
		friend.POST("/create-friends/:id", h.createFriends)
		friend.DELETE("/delete-friends/:id", h.deleteFriends)
	}

	newsfeed := router.Group("/newsfeed", h.userIdentity)
	{
		newsfeed.GET("/get-news/:id", h.getNews)
	}

	return router
}
