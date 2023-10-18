package main

import (
	"fio/internal/pkg/consumer"
	"fio/internal/pkg/kafka"
	"fio/internal/pkg/repository"
	"fio/internal/pkg/service"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("db.host"),
		Port:     os.Getenv("db.port"),
		Username: os.Getenv("db.username"),
		Password: os.Getenv("db.password"),
		DBName:   os.Getenv("db.dbname"),
		SSLMode:  os.Getenv("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	consumer := consumer.NewConsumer(services)
	kafka := kafka.NewKafka(consumer)

	err = kafka.Start()
	// if err != nil {
	// 	log.Fatalf("failed to start kafka: %s", err.Error())
	// }

}
