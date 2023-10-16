package main

import (
	"context"
	"fio/internal/consumer"
	"fio/internal/pkg/repository"
	"os"

	"log"

	kafkago "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

func main() {
	reader := consumer.NewKafkaReader()
	writer := consumer.NewKafkaWriter()

	ctx := context.Background()
	messages := make(chan kafkago.Message, 1000)
	messageCommitChan := make(chan kafkago.Message, 1000)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return reader.FetchMessageKafka(ctx, messages)
	})
	g.Go(func() error {
		return writer.WriteMessages(ctx, messages, messageCommitChan)
	})
	g.Go(func() error {
		return reader.CommitMessages(ctx, messageCommitChan)
	})
	err := g.Wait()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler.NewHandler(services)

}
