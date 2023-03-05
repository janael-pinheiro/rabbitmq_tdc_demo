package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/entities"
	amqp "github.com/rabbitmq/amqp091-go"
)

type amqpSubscriber struct {
	BrokerURL           string
	connection          connection
	notifyClosedChannel chan *amqp.Error
	queueName           string
	handler             MessageHandler
	consumerTag         string
}

func NewAMQPSubscriber(connection connection, queueName string, handler MessageHandler, consumerTag string) (*amqpSubscriber, error) {
	subscriber := new(amqpSubscriber)
	subscriber.connection = connection
	subscriber.queueName = queueName
	subscriber.handler = handler
	subscriber.consumerTag = consumerTag
	return subscriber, nil

}
func (as *amqpSubscriber) Subscribe() error {
	go as.reconnect()
	deliveries, err := as.connection.getChannel().Consume(as.queueName, as.consumerTag, AUTO_ACK, EXCLUSIVE, NO_LOCAL, NO_WAIT, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		fmt.Println("Start consuming!")
		for message := range deliveries {
			msg, err := as.unmarshal(message)
			// Simulates processing failure to send messages to the dead letter queue.
			if message.DeliveryTag%10 == 0 {
				message.Nack(MULTIPLE, REQUEUE)
				continue
			}
			if err == nil {
				if err := as.handler.Handle(msg); err != nil {
					message.Nack(MULTIPLE, REQUEUE)
				} else {
					message.Ack(MULTIPLE)
				}
			}
		}
	}()
	return nil
}

func (as *amqpSubscriber) reconnect() {
reconnection:
	for err := range as.connection.getClosedNotification() {
		fmt.Println("Reconnecting...")
		if err != nil {
			as.connection.connect()
			as.connection.createChannel()
			as.Subscribe()
			break reconnection
		}
	}
}

func (as *amqpSubscriber) Unsubscribe(id, topic string) error {
	return as.connection.getChannel().Cancel(as.consumerTag, NO_WAIT)
}

func (as *amqpSubscriber) Close() error {
	err := as.connection.getChannel().Close()
	if err != nil {
		return err
	}
	err = as.connection.getChannel().Close()
	return err
}

func (as *amqpSubscriber) unmarshal(message amqp.Delivery) (entities.Message, error) {
	msg := entities.Message{}
	err := json.Unmarshal(message.Body, &msg)
	if err != nil {
		return entities.Message{}, err
	}
	return msg, nil

}
