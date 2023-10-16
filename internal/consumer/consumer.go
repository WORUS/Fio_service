package consumer

import "fio/internal/pkg/service"

type Customer struct {
	services *service.Service
}

func NewCustomer(services *service.Service) *Customer {
	return &Customer{services: services}
}
