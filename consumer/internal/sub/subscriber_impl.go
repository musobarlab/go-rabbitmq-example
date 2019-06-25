package sub

import (
	"github.com/streadway/amqp"
)

//SubscriberImpl struct
type SubscriberImpl struct {
	c *amqp.Connection
}

//NewSubscriber constructor of SubscriberImpl
func NewSubscriber(address string) (*SubscriberImpl, error) {
	c, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}

	return &SubscriberImpl{c}, nil
}

//Subscribe function
func (s *SubscriberImpl) Subscribe(q string) (<-chan amqp.Delivery, func(), error) {
	ch, err := s.c.Channel()

	if err != nil {
		return nil, nil, err
	}

	// assert that the queue exists (creates a queue if it doesn't)
	//queue, err := ch.QueueDeclare(q, false, false, false, false, nil)

	// create a channel in go, through which incoming messages will be received
	c, err := ch.Consume(q, "", false, false, false, false, nil)

	// return the created channel
	return c, func() { ch.Close() }, nil
}
