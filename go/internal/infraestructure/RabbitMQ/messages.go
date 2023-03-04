package rabbitmq

import "github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/entities"

type MessageHandler interface {
	Handle(message entities.Message) error
}

type Publisher interface {
	Publish(topic string, message entities.Message) error
	Close() error
}

type Subscriber interface {
	Subscribe() error
	Unsubscribe(id, topic string) error
	Close() error
}
