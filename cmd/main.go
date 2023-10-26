package main

import (
	gql "fio/internal/pkg/graphql"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("error loading env variables: %s", err.Error())
	// }

	// db, err := repository.NewPostgresDB(repository.Config{
	// 	Host:     os.Getenv("db.host"),
	// 	Port:     os.Getenv("db.port"),
	// 	Username: os.Getenv("db.username"),
	// 	Password: os.Getenv("db.password"),
	// 	DBName:   os.Getenv("db.dbname"),
	// 	SSLMode:  os.Getenv("db.sslmode"),
	// })

	// if err != nil {
	// 	log.Fatalf("failed to initialize db: %s", err.Error())
	// }

	// repos := repository.NewRepository(db)
	// services := service.NewService(repos)
	// //consumer := consumer.NewConsumer(services)
	// handler := handler.NewHandler(services)
	// //kafka := kafka.NewKafka(consumer)
	// srv := new(server.Server)

	// func() {
	// 	if err := srv.Run(os.Getenv("port"), handler.InitRoutes()); err != nil {
	// 		log.Fatalf("error occured while running http server: %s", err.Error())
	// 	}
	// }()

	//err = kafka.Start()
	// if err != nil {
	// 	log.Fatalf("failed to start kafka consumer: %s", err.Error())
	// }
	schema, err := graphql.NewSchema(gql.DefineSchema())
	if err != nil {
		log.Panic("Error when creating the graphQL schema", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("Error when starting the http server", err)
	}
	log.Print("Server is started")

}
