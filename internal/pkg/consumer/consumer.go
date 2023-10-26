package consumer

import (
	"encoding/json"
	. "fio"
	"fio/internal/pkg/rest/service"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

type Record interface {
	CreateClient(client Client) error
}

type Consumer struct {
	Record
}

func NewConsumer(services *service.Service) *Consumer {
	return &Consumer{Record: NewRecordConsumer(services)}
}

func InitConsumer() {
	//TODO: move cunsomer methods from kafka
}

func (c *Consumer) CheckKafkaMessage(msg kafkago.Message) (Client, kafkago.Message, bool) {
	var client Client
	err := json.Unmarshal(msg.Value, &client)
	if err != nil {
		log.Printf("error decoding kafka message: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("kafka message: %q", msg.Value)
		note := []byte(` {"error":"incorrect format"}`)
		msg.Value = append(msg.Value, note...)
		return client, msg, false
	}
	if client.Name == "" || client.Surname == "" {
		note := []byte(` {"error":"no required field"}`)
		msg.Value = append(msg.Value, note...)
		return client, msg, false
	}
	client, err = c.Enrich(client)
	if err != nil {
		note := []byte(` {"error":"invalid name"}`)
		msg.Value = append(msg.Value, note...)
		return client, msg, false
	}
	return client, msg, true
}
