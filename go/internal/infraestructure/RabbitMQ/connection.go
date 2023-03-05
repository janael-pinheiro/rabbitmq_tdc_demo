package rabbitmq

import (
	"fmt"
	"os"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

type connection interface {
	getChannel() *amqp.Channel
	connect()
	getClosedNotification() chan *amqp.Error
	getPublishedNotification() chan amqp.Confirmation
	createChannel()
}

type amqpConnection struct {
	url                    string
	connection             *amqp.Connection
	channel                *amqp.Channel
	exchangeName           string
	exchangeType           string
	queueName              string
	routingKey             string
	notifyClosedChannel    chan *amqp.Error
	notifyPublishedChannel chan amqp.Confirmation
}

func NewConnection(url, exchangeName, exchangeType, queueName, routingKey string) connection {
	conn := new(amqpConnection)
	conn.url = url
	conn.exchangeName = exchangeName
	conn.exchangeType = exchangeType
	conn.queueName = queueName
	conn.routingKey = routingKey
	conn.connect()
	conn.createChannel()
	conn.declareExchange(exchangeName, exchangeType)
	conn.declareQueue()
	conn.bindQueue(queueName, exchangeName)
	return conn
}

func (ac *amqpConnection) createChannel() {
	channel, err := ac.connection.Channel()
	if err != nil {
		panic(err)
	}
	ac.channel = channel
	if err := ac.channel.Confirm(NO_WAIT); err == nil {
		ac.notifyPublishedChannel = ac.channel.NotifyPublish(make(chan amqp.Confirmation))
	}
	prefetchCount, err := strconv.Atoi(os.Getenv("PREFETCH_COUNT"))
	if err == nil {
		ac.channel.Qos(prefetchCount, PREFETCH_SIZE, GLOBAL_PREFETCH)
	}

}

func (ac *amqpConnection) getChannel() *amqp.Channel {
	return ac.channel
}

func (ac *amqpConnection) connect() {
	connection, err := amqp.Dial(ac.url)
	if err != nil {
		panic(err)
	}
	ac.connection = connection
	ac.notifyClosedChannel = ac.connection.NotifyClose(make(chan *amqp.Error))
}

func (ac *amqpConnection) declareQueue() {
	dead_letter_exchange := fmt.Sprintf("%s_dlx", ac.exchangeName)
	dead_letter_queue := fmt.Sprintf("%s_dlq", ac.queueName)
	arguments := amqp.Table{"x-dead-letter-exchange": dead_letter_exchange}
	ac.declareExchange(dead_letter_exchange, ac.exchangeType)
	_, queueErr := ac.channel.QueueDeclare(ac.queueName, DURABLE, AUTO_DELETE, EXCLUSIVE, NO_WAIT, arguments)
	_, dlQueueErr := ac.channel.QueueDeclare(dead_letter_queue, DURABLE, AUTO_DELETE, EXCLUSIVE, NO_WAIT, nil)
	if queueErr != nil || dlQueueErr != nil {
		panic(queueErr)
	}
	ac.bindQueue(dead_letter_queue, dead_letter_exchange)
}

func (ac *amqpConnection) declareExchange(exchangeName, exchangeType string) error {
	return ac.channel.ExchangeDeclare(exchangeName, exchangeType, DURABLE, AUTO_DELETE, INTERNAL, NO_WAIT, nil)
}

func (ac *amqpConnection) bindQueue(queueName, exchangeName string) {
	err := ac.channel.QueueBind(queueName, ac.routingKey, exchangeName, NO_WAIT, nil)
	if err != nil {
		panic(err)
	}
}

func (ac *amqpConnection) getClosedNotification() chan *amqp.Error {
	return ac.notifyClosedChannel
}

func (ac *amqpConnection) getPublishedNotification() chan amqp.Confirmation {
	return ac.notifyPublishedChannel
}
