package consumer

import (
	"context"
	"encoding/json"
	. "fio"
	"log"

	"github.com/pkg/errors"
	kafkago "github.com/segmentio/kafka-go"
)

type Reader struct {
	Reader *kafkago.Reader
}

func NewKafkaReader() *Reader {
	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "fio",
		GroupID: "group",
	})

	return &Reader{
		Reader: reader,
	}
}

func (k *Reader) FetchMessageKafka(ctx context.Context, messages chan kafkago.Message) error {
	for {
		message, err := k.Reader.FetchMessage(ctx)
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			var client Client
			if ok := checkFormat(&message, &client); !ok {
				messages <- message
				continue
			}
			log.Printf("\nname: %s, surname: %s", client.Name, client.Surname)
		}

	}
}

func (k *Reader) CommitMessages(ctx context.Context, messageCommitChan <-chan kafkago.Message) error {
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

func checkFormat(msg *kafkago.Message, c *Client) bool {
	err := json.Unmarshal(msg.Value, &c)
	if err != nil {
		log.Printf("\nerror decoding kafka message: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("kafka message: %q", msg.Value)
		note := []byte(` {"error":"incorrect format"}`)
		msg.Value = append(msg.Value, note...)
		return false
	}
	if c.Name == "" || c.Surname == "" {
		note := []byte(` {"error":"no required field"}`)
		msg.Value = append(msg.Value, note...)
		return false
	}
	if err := Enrich(c); err != nil {
		note := []byte(` {"error":"invalid name"}`)
		msg.Value = append(msg.Value, note...)
		return false
	}
	return true
}