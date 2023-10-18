package repository

import (
	. "fio"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RecordPostgres struct {
	db *sqlx.DB
}

func NewRecordPostgres(db *sqlx.DB) *RecordPostgres {
	return &RecordPostgres{db: db}
}

func (r *RecordPostgres) CreateClient(client Client) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, age, gender, country_id) values ($1, $2, $3, $4, $5) RETURNING id", clientTable)
	row := r.db.QueryRow(query, client.Name, client.Surname, client.Age, client.Gender, client.CountryId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
