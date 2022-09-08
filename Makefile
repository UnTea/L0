run:
	go run cmd/app/main.go

data:
	go run cmd/producer/nats_producer.go

.PHONY: run data