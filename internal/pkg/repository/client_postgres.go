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
	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, age, gender, country_id) values ($1, $2, $3, $4, $5, $6) RETURNING id", clientsTable)
	row := r.db.QueryRow(query, client.Name, client.Surname, client.Patronymic, client.Age, client.Gender, client.CountryId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RecordPostgres) GetClients(username, password string) (Client, error) {
	var user Client
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", clientsTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
