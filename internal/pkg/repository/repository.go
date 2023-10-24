package repository

import (
	. "fio"

	"github.com/jmoiron/sqlx"
)

type Record interface {
	CreateClient(client Client) (int, error)
	GetClientsByFilter(sql string) ([]Client, error)
}

type Repository struct {
	Record
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Record: NewRecordPostgres(db),
	}

}
