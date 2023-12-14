package service

import (
	"context"
	"fmt"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/entity"
)

func (s *service) AddUser(ctx context.Context, user entity.UserExtend) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	errAddUser := s.usersRepo.Add(ctx, user)
	if errAddUser != nil {
		return fmt.Errorf("[service.AddUser]->%w", errAddUser)
	}

	return nil
}

func (s *service) GetUser(ctx context.Context, username string) (entity.UserExtend, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, errGet := s.usersRepo.Get(ctx, username)
	if errGet != nil {
		return entity.UserExtend{}, fmt.Errorf("[service.GetUser]->%w", errGet)
	}
	return user, nil
}

func (s *service) ContainsUser(ctx context.Context, username string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	contains, errContains := s.usersRepo.Contains(ctx, username)
	if errContains != nil {
		return false, fmt.Errorf("[service.ContainsUser]->%w", errContains)
	}
	return contains, nil
}
