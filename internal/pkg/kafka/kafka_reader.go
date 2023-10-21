package kafka

import (
	"context"

	"log"

	"github.com/pkg/errors"
	kafkago "github.com/segmentio/kafka-go"
)

func NewKafkaReader() *kafkago.Reader {
	return kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "fio",
		GroupID: "group",
	})

}

func (k *Kafka) FetchMessageKafka(ctx context.Context, messages chan kafkago.Message) error {
	for {
		message, err := k.Reader.FetchMessage(ctx)
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			client, message, ok := k.consumer.CheckKafkaMessage(message)
			if !ok {
				messages <- message
				continue
			}
			log.Printf("name: %s, surname: %s, patronymic: %s, age: %d, gender: %s", client.Name, client.Surname, client.Patronymic, client.Age, client.Gender)
			k.consumer.CreateClient(client)
		}

	}
}

func (k *Kafka) CommitMessages(ctx context.Context, messageCommitChan <-chan kafkago.Message) error {
	for {
		select {
		case <-ctx.Done():
		case msg := <-messageCommitChan:
			err := k.Reader.CommitMessages(ctx, msg)
			if err != nil {
				return errors.Wrap(err, "Reader.CommitMessages")
			}
			//log.Printf("commited an msg : %v \n", string(msg.Value))
		}
	}
}
