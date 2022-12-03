package service

import (
	"github.com/dhevve/blog/internal/model"
	"github.com/dhevve/blog/internal/repository"
)

type FriendServce struct {
	repo repository.Friend
}

func NewFriendServce(repo repository.Friend) *FriendServce {
	return &FriendServce{repo: repo}
}

func (s *FriendServce) GetFriends(id int) []model.User {
	return s.repo.GetFriends(id)
}

func (s *FriendServce) CreateFriends(myId, friendId int) error {
	return s.repo.CreateFriends(myId, friendId)
}

func (s *FriendServce) DeleteFriend(myId, friendId int) error {
	return s.repo.DeleteFriend(myId, friendId)
}
