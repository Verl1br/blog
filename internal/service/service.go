package service

import (
	"github.com/dhevve/blog/internal/repository"
)

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
