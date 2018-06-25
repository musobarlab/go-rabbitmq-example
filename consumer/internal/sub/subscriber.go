package sub

import(
	"github.com/streadway/amqp"
)

//Subscriber interface
type Subscriber interface {
	Subscribe(string) (<-chan amqp.Delivery, func(), error)
}