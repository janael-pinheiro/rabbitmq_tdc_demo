from dataclasses import dataclass
from logging import Logger
from pika import URLParameters
from pika.exceptions import ConnectionClosedByBroker
from pika.channel import Channel
from tenacity import retry
from tenacity.retry import retry_if_exception_type
from tenacity.wait import wait_exponential
from os import environ

from rabbitmq.amqp.connection import (
    AMQPConnection,
    AMQPChannel)


def callback(channel, method, properties, body):
    print(body)


@dataclass
class Subscriber:
    channel: Channel
    queue_name: str
    routing_key: str
    logger: Logger
    
    @retry(
        retry=retry_if_exception_type(ConnectionClosedByBroker),
        wait=wait_exponential(multiplier=1, min=4, max=10))
    def subscribe(self):
        if self.channel.connection.is_closed:
            self.logger.info("Subscriber connection closed! Reconnecting...")
            connection = AMQPConnection(URLParameters(environ.get("AMQP_URL"))).create()
            self.channel = AMQPChannel(connection=connection).create()
        self.__start()

    def unsubscribe(self):
        self.channel.cancel()

    def __start(self):
        self.channel.basic_consume(
            queue=self.queue_name,
            on_message_callback=callback,
            auto_ack=True)
        try:
            self.channel.start_consuming()
        except KeyboardInterrupt:
            self.channel.stop_consuming()
