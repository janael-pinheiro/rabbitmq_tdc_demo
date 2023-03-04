package main

import (
	"github.com/google/uuid"
	rabbitmq "github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/RabbitMQ"
	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/common"
)

func main() {
	queueName := "test"
	exchangeName := "demo"
	exchangeType := "fanout"
	routingKey := "123A"
	consumerTag := uuid.NewString()
	broker_configuration := common.LoadBrokerConfiguration()
	connection := rabbitmq.NewConnection(broker_configuration.URL, exchangeName, exchangeType, queueName, routingKey)
	messageHandler := new(rabbitmq.AMQPMessageHandler)
	subscriber, err := rabbitmq.NewAMQPSubscriber(connection, queueName, messageHandler, consumerTag)
	if err != nil {
		panic(err)
	}
	subscriber.Subscribe()
	waitForever := make(chan struct{})
	<-waitForever
}
