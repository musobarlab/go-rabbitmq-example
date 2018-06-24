package pub

import (
	"fmt"

	"github.com/streadway/amqp"
)

//PublisherImpl struct
type PublisherImpl struct {
	c *amqp.Connection
}

//NewPublisher constructor of PublisherImpl
func NewPublisher(address string) (*PublisherImpl, error) {
	c, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}

	return &PublisherImpl{c}, nil
}

//Publish function
func (p *PublisherImpl) Publish(q string, message []byte) error {
	ch, err := p.c.Channel()

	if err != nil {
		return err
	}

	defer ch.Close()

	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType: "application/json",
		Body: message,
	}

	if err := ch.Publish("", q, false, false, payload); err != nil {
		return fmt.Errorf("[Publish] error %v", err)
	}

	return nil
}