package service

import (
	"context"

	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
)

type FriendServce struct {
	repo repository.Friend
}

func NewFriendServce(repo repository.Friend) *FriendServce {
	return &FriendServce{repo: repo}
}

func (s *FriendServce) GetFriends(id int, ctx context.Context) []model.User {
	return s.repo.GetFriends(id, ctx)
}

func (s *FriendServce) CreateFriends(myId, friendId int, ctx context.Context) error {
	return s.repo.CreateFriends(myId, friendId, ctx)
}

func (s *FriendServce) DeleteFriend(myId, friendId int, ctx context.Context) error {
	return s.repo.DeleteFriend(myId, friendId, ctx)
}
