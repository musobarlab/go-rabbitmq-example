package handler

import (
	"fmt"
	"encoding/json"
	"github.com/wuriyanto48/go-rabbitmq-example/consumer/internal/sub"
)

//WorkerHandler struct
type WorkerHandler struct {
	q string
	subscriber sub.Subscriber
}

//NewWorkerHandler WorkerHandler's constructor
func NewWorkerHandler(q string, subscriber sub.Subscriber) *WorkerHandler {
	return &WorkerHandler{q: q, subscriber: subscriber}
}

//Pool function
func (h *WorkerHandler) Pool(){
	messages, close, err := h.subscriber.Subscribe(h.q)
	if err != nil {
		panic(err)
	}

	defer close()

	forever := make(chan bool)

	go func(){
		//loop over chan messages
		for msg := range messages {
			//receiver message
			var message sub.Message
			_ = json.Unmarshal(msg.Body, &message)

			fmt.Println(message)

			msg.Ack(false)
		}
	}()
	<-forever
}