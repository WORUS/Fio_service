package repository

import (
	. "fio"
	"fmt"
	"log"
	"strings"

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

func (r *RecordPostgres) GetClientsByFilter(sql string, page int) ([]Client, error) {
	var clientSQL []ClientSQL
	limit := 2
	offset := limit * (page - 1)

	query := fmt.Sprintf("SELECT * FROM %s%s LIMIT $2 OFFSET $1", clientsTable, sql)
	log.Println(query)
	err := r.db.Select(&clientSQL, query, offset, limit)
	if err != nil {
		return nil, err
	}

	clients := make([]Client, len(clientSQL))
	for i := range clientSQL {
		clients[i].ID = int(clientSQL[i].ID.Int64)
		clients[i].Name = clientSQL[i].Name.String
		clients[i].Surname = clientSQL[i].Surname.String
		clients[i].Patronymic = clientSQL[i].Patronymic.String
		clients[i].Age = int(clientSQL[i].Age.Int16)
		clients[i].Gender = clientSQL[i].Gender.String
		clients[i].CountryId = clientSQL[i].CountryId.String
	}

	return clients, nil
}

func (r *RecordPostgres) UpdateClientRecord(id int, client ClientUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if client.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *client.Name)
		argId++
	}
	if client.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *client.Surname)
		argId++
	}
	if client.Patronymic != nil {
		setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
		args = append(args, *client.Patronymic)
		argId++
	}
	if client.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *client.Age)
		argId++
	}
	if client.Gender != nil {
		setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
		args = append(args, *client.Gender)
		argId++
	}
	if client.CountryId != nil {
		setValues = append(setValues, fmt.Sprintf("country_id=$%d", argId))
		args = append(args, *client.CountryId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = %d`, clientsTable, setQuery, id)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *RecordPostgres) DeleteClientById(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, clientsTable)
	_, err := r.db.Exec(query, id)
	return err
}
