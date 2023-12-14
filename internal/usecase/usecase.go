package usecase

import (
	"sync"

	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/interfaces"
)

type usecase struct {
	mu      *sync.RWMutex
	service interfaces.IService
	genID   interfaces.IGeneratorID
}

var _ interfaces.IUseCase = (*usecase)(nil)

func NewUseCase(s interfaces.IService, g interfaces.IGeneratorID) *usecase {
	return &usecase{
		mu:      &sync.RWMutex{},
		service: s,
		genID:   g,
	}
}
