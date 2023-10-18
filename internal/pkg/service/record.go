package service

import (
	. "fio"
	"fio/internal/pkg/repository"
)

type RecordService struct {
	repo repository.Record
}

func NewRecordService(repo repository.Record) *RecordService {
	return &RecordService{repo: repo}
}

func (s *RecordService) CreateClient(client Client) (int, error) {
	return s.repo.CreateClient(client)
}
