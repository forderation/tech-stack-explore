package nsqt

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

type messageHandler struct{}

func ReceiveTopic() {
	config := nsq.NewConfig()

	// giving up process 10x
	config.MaxAttempts = 10

	// maximum number messages to allow in flight
	config.MaxInFlight = 5

	// Maximum duration when REQueueing
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0

	// Init topic name and channel
	topic := "Topic_Example"
	channel := "Channel_Example"

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Handler for messages received by this Consumer.
	consumer.AddHandler(&messageHandler{})

	//Use nsqlookupd to find nsqd instances
	consumer.ConnectToNSQLookupd("127.0.0.1:4161")

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumer.Stop()
}

func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	var request Message
	if err := json.Unmarshal(m.Body, &request); err != nil {
		log.Println("Error when unmarshall ", err)
		return err
	}

	log.Println("Message")
	log.Println("-----------")
	log.Println("Name: ", request.Name)
	log.Println("Content: ", request.Content)
	log.Println("Timestamp: ", request.Timestamp)
	log.Println("-----------")
	return nil
}
