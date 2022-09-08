package main

import (
	"encoding/json"
	"fmt"
	"github.com/UnTea/L0/internal/model"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
)

const (
	clientID  = "publisher"
	clusterID = "test-cluster"
	channel   = "session"
)

func main() {

	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err := sc.Close()
		if err != nil {
			fmt.Printf("Error occurred while closing NATS-streaming connection: %v", err)
		}
	}()

	jsonFile, err := os.Open("model.json")
	if err != nil {
		log.Printf("Error occurred while openning JSON file: %v", err)
	}

	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Printf("Error occurred while closing JSON file: %v", err)
		}
	}()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data []model.Data

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Printf("Error occurred while parsing JSON file %v", err)
	}

	data = append(data, model.Data{})

	for _, value := range data {
		bytesValue, _ := json.Marshal(value)

		err = sc.Publish(channel, bytesValue)
		if err != nil {
			fmt.Printf("Error occurred while publishing data into cluster: %v", err)
		}
	}

	err = sc.Publish(channel, []byte("Some wrong data"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Out of 6 messages sent: 4 was correct, 2 was incorrect")
}
