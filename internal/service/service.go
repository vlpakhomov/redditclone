package service

import (
	"sync"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

type service struct {
	mu        *sync.RWMutex
	postsRepo interfaces.IPostRepository
	usersRepo interfaces.IUserRepository
}

var _ interfaces.IService = (*service)(nil)

func NewService(posts interfaces.IPostRepository, users interfaces.IUserRepository) *service {
	return &service{
		mu:        &sync.RWMutex{},
		postsRepo: posts,
		usersRepo: users,
	}
}
