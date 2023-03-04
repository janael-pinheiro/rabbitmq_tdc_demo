package common

import (
	"os"

	"github.com/janael-pinheiro/rabbitmq_tdc_demo/internal/infraestructure/dto"
)

func LoadBrokerConfiguration() *dto.BrokerConfigurationDTO {
	configuration := new(dto.BrokerConfigurationDTO)
	configuration.URL = os.Getenv("AMQP_URL")
	return configuration
}
