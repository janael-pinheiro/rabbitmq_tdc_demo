from dataclasses import dataclass
from logging import Logger
from os import environ
from typing import Dict, Any
from pika import URLParameters, BasicProperties, spec
from pika.channel import Channel
from pika.exceptions import ConnectionClosedByBroker
from tenacity import retry
from tenacity.retry import retry_if_exception_type
from tenacity.wait import wait_exponential

from rabbitmq.amqp.connection import AMQPChannel, AMQPConnection
from rabbitmq.amqp.constants import AUTO_ACK, MULTIPLE, REQUEUE


def callback(
        channel: Channel,
        method: spec.Basic.Deliver,
        properties: BasicProperties,
        body: Dict[str, Any]):
    try:
        print(body)
    except Exception as _:
        channel.basic_nack(
            delivery_tag=method.delivery_tag,
            multiple=MULTIPLE,
            requeue=REQUEUE)
    else:
        channel.basic_ack(delivery_tag=method.delivery_tag, multiple=MULTIPLE)


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
            auto_ack=AUTO_ACK)
        try:
            self.channel.start_consuming()
        except KeyboardInterrupt:
            self.channel.stop_consuming()
