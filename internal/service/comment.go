package service

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
)

type CommentService struct {
	repo repository.Сomment
}

func NewCommentService(repo repository.Сomment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment model.Comment) (int, error) {
	return s.repo.CreateComment(comment)
}

func (s *CommentService) GetComment(commentId int) (model.Comment, error) {
	return s.repo.GetComment(commentId)
}

func (s *CommentService) GetComments() ([]model.Comment, error) {
	return s.repo.GetComments()
}

func (s *CommentService) DeleteComment(commentId int) error {
	return s.repo.DeleteComment(commentId)
}

func (s *CommentService) UpdateComment(commentId int, input model.UpdateComment) error {
	return s.repo.UpdateComment(commentId, input)
}
