package service

import (
	. "fio"
	"fio/internal/pkg/repository"
)

type Record interface {
	CreateClient(Client Client) (int, error)
	GetClient()
}

type Service struct {
	Record
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Record: NewClientService(repos.Client),
	}
}
