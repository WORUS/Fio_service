package repository

import (
	. "fio"

	"github.com/jmoiron/sqlx"
)

type Record interface {
	CreateClient(client Client) (int, error)
	GetClientByName()
}

type Repository struct {
	Record
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Record: NewClientPostgres(db),
	}

}
