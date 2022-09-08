package main

import (
	"context"
	"fmt"
	"github.com/UnTea/L0/internal/api"
	"github.com/UnTea/L0/internal/config"
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
	ctx := context.Background()
	ch := make(chan model.Data, 10)

	cfg, err := config.ReadConfigYML("config.yml")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while reading configuration file: %v", err))
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := helpers.NewPostgres(ctx, dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while creating connection with postgresql: %v", err))
	}

	defer db.Close()

	err = migrations.MakeMigrations(dsn, cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while migrationing database: %v", err))
	}

	sc, err := stan.Connect(cfg.Nats.ClusterID, cfg.Nats.ClientID)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while creating connection with NATS-streaming: %v", err))
	}

	defer helpers.Closer(sc)

	sub, err := nats.NewSubscription(sc, cfg, ch)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while creating new subscription: %v", err))
	}

	defer func() {
		err := sub.Unsubscribe()
		if err != nil {
			fmt.Printf("Error occurred while remove subscription: %v", err)
		}
	}()

	defer close(ch)

	router := httprouter.New()
	handler := api.NewHandler(db, ch)

	handler.Register(router)

	address := fmt.Sprintf("%s:%s", cfg.HttpServer.IP, cfg.HttpServer.Port)
	log.Printf("Server is listening on: %s\n", address)

	err = http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error occurred while creating tcp connection: %v", err))
	}
}
