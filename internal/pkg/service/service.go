package service

import (
	. "fio"
	"fio/internal/pkg/repository"
)

type Record interface {
	CreateClient(client Client) (int, error)
	GetClientsByFilter(filter ClientFilter) ([]Client, error)
}

type Service struct {
	Record
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Record: NewRecordService(repos.Record),
	}
}
