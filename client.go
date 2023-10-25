package fio

import "database/sql"

type Client struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
	CountryId  string `json:"country_id" binding:"required"`
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
	Name       sql.NullString `db:"name"`
	Surname    sql.NullString `db:"surname"`
	Patronymic sql.NullString `db:"patronymic"`
	Age        sql.NullInt16  `db:"age"`
	Gender     sql.NullString `db:"gender"`
	CountryId  sql.NullString `db:"country_id"`
}

type ClientUpdate struct {
	ID         int     `json:"id"`
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`
	Age        *int    `json:"age"`
	Gender     *string `json:"gender"`
	CountryId  *string `json:"country_id"`
}
