package nsqt

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func SendTopic() {
	config := nsq.NewConfig()

	// connect to daemon
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	// init topic and message
	topic := "Topic_Example"
	msg := Message{
		Name:      "Message Name Example",
		Content:   "Message Content Example",
		Timestamp: time.Now().String(),
	}

	// convert message as byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	err = producer.Publish(topic, payload)
	if err != nil {
		panic(err)
	}
}
