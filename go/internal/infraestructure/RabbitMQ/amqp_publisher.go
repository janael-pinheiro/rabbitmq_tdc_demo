package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/entities"
	amqp "github.com/rabbitmq/amqp091-go"
)

type amqpPublisher struct {
	BrokerURL             string
	connection            connection
	exchange              string
	notificationChannel   chan *amqp.Error
	notificationPublished chan amqp.Confirmation
}

func NewAMQPPublisher(brokerURL, exchangeName string, connection connection) (*amqpPublisher, error) {
	publisher := new(amqpPublisher)
	publisher.connection = connection
	publisher.exchange = exchangeName
	publisher.notificationChannel = publisher.connection.getClosedNotification()
	publisher.notificationPublished = publisher.connection.getPublishedNotification()
	return publisher, nil
}

func (pub *amqpPublisher) Publish(message entities.Message) error {
	var err error
	data, dataErr := json.Marshal(message)
	if dataErr != nil {
		return dataErr
	}
	//Checks if the connection is active.
	select {
	case err := <-pub.notificationChannel:
		if err != nil {
			pub.reconnect()
		}
	default:
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = pub.connection.getChannel().PublishWithContext(ctx, pub.exchange, "", MANDATORY, IMMEDIATE, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data})
	pub.getPublishedConfirmation()
	return err
}

func (pub *amqpPublisher) reconnect() {
	fmt.Println("Reconnecting...")
	pub.connection.connect()
	pub.connection.createChannel()
	pub.notificationChannel = pub.connection.getClosedNotification()
	pub.notificationPublished = pub.connection.getPublishedNotification()
}

func (pub *amqpPublisher) getPublishedConfirmation() {
	notification := <-pub.notificationPublished
	fmt.Printf("Delivery tag: %d -- published: %t\n", notification.DeliveryTag, notification.Ack)
}
