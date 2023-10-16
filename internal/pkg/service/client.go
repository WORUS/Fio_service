package service

import (
	. "fio"
	"fio/internal/pkg/repository"
)

type ClientService struct {
	repo repository.Record
}

func NewClientService(repo repository.Record) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) CreateClient(client Client) (int, error) {
	return s.repo.CreateClient(client)
}
