package service

import (
	. "fio"
	"fio/internal/pkg/repository"
	"fmt"
	"strconv"
	"strings"
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

func (s *RecordService) GetClientsByFilter(filter ClientFilter) ([]Client, error) {
	var query []string
	sql := ""

	if filter.Name != nil {
		names := strings.Join(filter.Name, "','")
		query = append(query, fmt.Sprintf("name IN ('%s')", names))
	}
	if filter.Surname != nil {
		surnames := strings.Join(filter.Surname, "','")
		query = append(query, fmt.Sprintf("surnames IN ('%s')", surnames))
	}
	if filter.Patronymic != nil {
		patronymics := strings.Join(filter.Patronymic, "','")
		query = append(query, fmt.Sprintf("patronymic IN ('%s')", patronymics))
	}
	if filter.Age != nil {
		ages := strconv.Itoa(filter.Age[0]) + " AND " + strconv.Itoa(filter.Age[1])
		query = append(query, fmt.Sprintf("age BETWEEN %s", ages))
	}
	if filter.Gender != nil {
		genders := strings.Join(filter.Gender, "','")
		query = append(query, fmt.Sprintf("gender IN ('%s')", genders))
	}
	if filter.CountryId != nil {
		countryIds := strings.Join(filter.CountryId, "','")
		query = append(query, fmt.Sprintf("country_id IN ('%s')", countryIds))
	}

	for i := range query {
		sql += query[i]
		if i != len(query)-1 {
			sql += " AND "
		}
	}
	return s.repo.GetClientsByFilter(sql)
}
