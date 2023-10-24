package repository

import (
	. "fio"
	"fmt"
	"log"

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

func (r *RecordPostgres) GetClientsByFilter(sql string) ([]Client, error) {
	var clientSQL []ClientSQL

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", clientsTable, sql)
	log.Println(query)
	err := r.db.Select(&clientSQL, query)
	if err != nil {
		return nil, err
	}

	clients := make([]Client, len(clientSQL))
	for i := range clientSQL {
		clients[i].Name = clientSQL[i].Name.String
		clients[i].Surname = clientSQL[i].Surname.String
		clients[i].Patronymic = clientSQL[i].Patronymic.String
		clients[i].Age = int(clientSQL[i].Age.Int16)
		clients[i].Gender = clientSQL[i].Gender.String
		clients[i].CountryId = clientSQL[i].CountryId.String
	}

	return clients, nil
}
