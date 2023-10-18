package kafka

import (
	"context"
	"fio/internal/pkg/consumer"
	"log"

	kafkago "github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"
)

type Kafka struct {
	consumer *consumer.Consumer
	Reader   *kafkago.Reader
	Writer   *kafkago.Writer
}

func NewKafka(consumer *consumer.Consumer) *Kafka {
	return &Kafka{
		consumer: consumer,
		Reader:   NewKafkaReader(),
		Writer:   NewKafkaWriter()}
}

// func (k *Kafka) InitClient(client Client) {
// 	k.consumer.CreateClient(client)
// }

func (k *Kafka) Start() error {

	ctx := context.Background()
	messages := make(chan kafkago.Message, 1000)
	messageCommitChan := make(chan kafkago.Message, 1000)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return k.FetchMessageKafka(ctx, messages)
	})
	g.Go(func() error {
		return k.WriteMessages(ctx, messages, messageCommitChan)
	})
	g.Go(func() error {
		return k.CommitMessages(ctx, messageCommitChan)
	})
	err := g.Wait()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
