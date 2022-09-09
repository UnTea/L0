package nats

import (
	"encoding/json"
	"github.com/UnTea/L0/internal/config"
	"github.com/UnTea/L0/internal/model"
	"github.com/nats-io/stan.go"
	"log"
)

var channel chan model.Data

func handle(message *stan.Msg) {
	var data model.Data

	err := json.Unmarshal(message.Data, &data)
	if err != nil {
		log.Printf("Error occurred while decoding data from nats channel: %v ", err)
		message.Ack()

		return
	}

	channel <- data

	if err := message.Ack(); err != nil {
		log.Printf("failed tp ACK msg: %d", message.Sequence)

		return
	}
}

func NewSubscription(connection stan.Conn, cfg config.Config, ch chan model.Data) (stan.Subscription, error) {
	channel = ch

	sub, err := connection.Subscribe(
		cfg.Nats.Channel,
		handle,
		stan.DurableName("last-position"),
		stan.SetManualAckMode(),
	)
	if err != nil {
		return nil, err
	}

	return sub, nil
}
