package rabbitmq

import (
	"fmt"

	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/entities"
)

type AMQPMessageHandler struct{}

func (mh *AMQPMessageHandler) Handle(message entities.Message) error {
	fmt.Println("Message: ", message.Content)
	return nil
}
