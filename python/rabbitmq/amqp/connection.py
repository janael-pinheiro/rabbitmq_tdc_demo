from dataclasses import dataclass
from typing import Union

from pika import BlockingConnection
from pika.channel import Channel
from pika.connection import Parameters

from rabbitmq.amqp.constants import (AUTO_DELETE, DURABLE, EXCLUSIVE,
                                     GLOBAL_PREFETCH, PREFETCH_COUNT,
                                     PREFETCH_SIZE)


@dataclass
class AMQPConnection:
    parameters: Parameters

    def create(self) -> BlockingConnection:
        return BlockingConnection(self.parameters)

@dataclass
class AMQPChannel:
    connection: BlockingConnection

    def create(self) -> Channel:
        channel = self.connection.channel()
        channel.confirm_delivery()
        channel.basic_qos(
            prefetch_size=PREFETCH_SIZE,
            prefetch_count=PREFETCH_COUNT,
            global_qos=GLOBAL_PREFETCH)
        return channel


@dataclass
class AMQPExchange:
    exchange_name: str
    exchange_type: str
    channel: Channel

    def declare(self) -> None:
        self.channel.exchange_declare(
            exchange=self.exchange_name,
            exchange_type=self.exchange_type,
            durable=DURABLE,
            auto_delete=AUTO_DELETE)


@dataclass
class AMQPQueue:
    channel: Channel
    name: str
    dead_letter_exchange: Union[str, None] = None

    def declare(self) -> None:
        arguments = None
        if self.dead_letter_exchange:
            arguments = arguments={"x-dead-letter-exchange": self.dead_letter_exchange}
        self.channel.queue_declare(
            queue=self.name,
            durable=DURABLE,
            exclusive=EXCLUSIVE,
            auto_delete=AUTO_DELETE,
            arguments=arguments
            )

    def bind(self, exchange_name: str, routing_key: str) -> None:
        self.channel.queue_bind(
            queue=self.name,
            exchange=exchange_name,
            routing_key=routing_key)
