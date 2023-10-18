package consumer

import (
	. "fio"
	"fio/internal/pkg/service"
	"log"
)

type RecordConsumer struct {
	services service.Record
}

func NewRecordConsumer(services service.Record) *RecordConsumer {
	return &RecordConsumer{services: services}
}

func (rc *RecordConsumer) CreateClient(client Client) error {
	id, err := rc.services.CreateClient(client)
	if err != nil {
		log.Printf("error occured %v", err)
		return err
	}

	log.Printf("create user with id = %d", id)
	return nil
}
