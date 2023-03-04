package main

import (
	"fmt"
	"time"

	rabbitmq "github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/RabbitMQ"
	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/common"
	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/entities"
)

func main() {
	queueName := "go"
	exchangeName := "demo"
	exchangeType := "fanout"
	routingKey := "123A"
	broker_configuration := common.LoadBrokerConfiguration()
	connection := rabbitmq.NewConnection(broker_configuration.URL, exchangeName, exchangeType, queueName, routingKey)
	publisher, err := rabbitmq.NewAMQPPublisher(broker_configuration.URL, exchangeName, connection)
	if err != nil {
		panic(err)
	}
	message := entities.Message{Content: "Hello, world!"}
	fmt.Println("Application started")
	for {
		time.Sleep(time.Millisecond * 100)
		err = publisher.Publish(message)
		if err != nil {
			panic(err)
		}
	}
}
