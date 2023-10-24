package fio

import "database/sql"

type Client struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name" binding:"required"`
	Surname    string `json:"surname" db:"surname" binding:"required"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Age        int    `json:"age" db:"age"`
	Gender     string `json:"gender" db:"gender"`
	CountryId  string `json:"country_id" db:"country_id"`
}

type ClientFilter struct {
	Name       []string
	Surname    []string
	Patronymic []string
	Age        []int
	Gender     []string
	CountryId  []string
}

type ClientSQL struct {
	ID         sql.NullInt64  `db:"id"`
	Name       sql.NullString `db:"name" binding:"required"`
	Surname    sql.NullString `db:"surname" binding:"required"`
	Patronymic sql.NullString `db:"patronymic"`
	Age        sql.NullInt16  `db:"age"`
	Gender     sql.NullString `db:"gender"`
	CountryId  sql.NullString `db:"country_id"`
}
