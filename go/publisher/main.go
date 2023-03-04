package main

import (
	"fmt"
	"time"

	rabbitmq "github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/RabbitMQ"
	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/common"
	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/entities"
)

func main() {
	broker_configuration := common.LoadBrokerConfiguration()
	connection := rabbitmq.NewConnection(broker_configuration.URL, "demo", "fanout", "test", "123A")
	publisher, err := rabbitmq.NewAMQPPublisher(broker_configuration.URL, "demo", connection)
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
