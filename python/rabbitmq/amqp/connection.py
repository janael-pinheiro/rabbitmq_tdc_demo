from dataclasses import dataclass

from pika import BlockingConnection
from pika.channel import Channel
from pika.connection import Parameters
from rabbitmq.amqp.constants import DURABLE, AUTO_DELETE, EXCLUSIVE

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

    def declare(self) -> None:
        self.channel.queue_declare(
            queue=self.name,
            durable=DURABLE,
            exclusive=EXCLUSIVE,
            auto_delete=AUTO_DELETE)

    def bind(self, exchange_name: str, routing_key: str) -> None:
        self.channel.queue_bind(
            queue=self.name,
            exchange=exchange_name,
            routing_key=routing_key)
