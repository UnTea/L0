package main

import (
	ctx "context"
	"fmt"
	"github.com/UnTea/L0/internal/api"
	cfg "github.com/UnTea/L0/internal/config"
	"github.com/UnTea/L0/internal/migrations"
	"github.com/UnTea/L0/internal/model"
	"github.com/UnTea/L0/internal/nats"
	"github.com/UnTea/L0/pkg/helpers"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
)

func main() {
	context := ctx.Background()
	channel := make(chan model.Data, 10)

	config, err := cfg.ReadConfigYML("config.yml")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while reading configuration file: %v", err))
	}

	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.SslMode,
	)

	database, err := helpers.NewPostgres(context, connectionString)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while creating connection with postgresql: %v", err))
	}

	defer database.Close()

	err = migrations.MakeMigrations(connectionString, config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while migrationing database: %v", err))
	}

	stanConnection, err := stan.Connect(config.Nats.ClusterID, config.Nats.ClientID)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while creating connection with NATS-streaming: %v", err))
	}

	defer helpers.Closer(stanConnection)

	subscription, err := nats.NewSubscription(stanConnection, config, channel)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while creating new subscription: %v", err))
	}

	defer func() {
		err := subscription.Unsubscribe()
		if err != nil {
			fmt.Printf("Error occurred while remove subscription: %v", err)
		}
	}()

	defer close(channel)

	router := httprouter.New()
	handler := api.NewHandler(database, channel)

	handler.Register(router)

	address := fmt.Sprintf("%s:%s", config.HttpServer.IP, config.HttpServer.Port)
	log.Printf("Server is listening on: %s\n", address)

	err = http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while creating tcp connection: %v", err))
	}
}
