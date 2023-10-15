package fio

import (
	"fio/pkg/consumer"
)

type User struct {
	cons        consumer.User
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}
