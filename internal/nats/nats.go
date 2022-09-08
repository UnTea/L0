package nats

import (
	"encoding/json"
	"github.com/UnTea/L0/internal/config"
	"github.com/UnTea/L0/internal/model"
	"github.com/nats-io/stan.go"
	"log"
)

var ch chan model.Data

func handle(msg *stan.Msg) {
	var data model.Data

	err := json.Unmarshal(msg.Data, &data)
	if err != nil {
		log.Printf("error occurred while decoding data from nats channel: %v ", err)
		msg.Ack()

		return
	}

	// sending data to service
	ch <- data

	if err := msg.Ack(); err != nil {
		log.Printf("failed tp ACK msg: %d", msg.Sequence)

		return
	}
}

// NewSubscription - creates new nats subscription
func NewSubscription(conn stan.Conn, cfg config.Config, c chan model.Data) (stan.Subscription, error) {
	ch = c
	sub, err := conn.Subscribe(
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
