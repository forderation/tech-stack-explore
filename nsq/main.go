package main

import (
	"fmt"
	"nsq-learn/nsqt"
	"time"
)

func main() {
	nsqt.SendTopic()
	fmt.Println("Wait for 2 second")
	time.Sleep(2 * time.Second)
	nsqt.ReceiveTopic()
}
