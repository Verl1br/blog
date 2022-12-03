package service

import (
	"github.com/dhevve/blog/internal/repository"
	"github.com/gin-gonic/gin"
)

type PhotoService struct {
	repo repository.Photo
}

func NewPhotoService(repo repository.Photo) *PhotoService {
	return &PhotoService{repo: repo}
}

func (s *PhotoService) Upload(c *gin.Context, postId int) (int, error) {
	fullFileName, err := saveFile(c)
	if err != nil {
		return 0, err
	}
	return s.repo.Upload(postId, fullFileName)
}

func (s *PhotoService) DeletePhoto(photoId int) error {
	return s.repo.DeletePhoto(photoId)
}
