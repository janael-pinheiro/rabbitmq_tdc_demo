from pika.channel import Channel
from pika import BlockingConnection
from pika import URLParameters
from os import environ

from rabbitmq.amqp.connection import (
    AMQPConnection,
    AMQPChannel,
    AMQPExchange,
    AMQPQueue)


def create_connection() -> BlockingConnection:
    parameters = URLParameters(environ.get("AMQP_URL"))
    return AMQPConnection(parameters=parameters).create()
    

def create_channel(connection: BlockingConnection) -> Channel:
    return AMQPChannel(connection=connection).create()


def create_exchange(name: str, exchange_type: str,  channel: Channel) -> None:
    AMQPExchange(
        channel=channel,
        exchange_name=name,
        exchange_type=exchange_type).declare()


def create_queue(
        name: str,
        channel: Channel,
        exchange_name: str,
        routing_key: str,
        dead_letter_exchange: str = None) -> None:
    queue = AMQPQueue(
        name=name,
        channel=channel,
        dead_letter_exchange=dead_letter_exchange)
    queue.declare()
    queue.bind(
        exchange_name=exchange_name,
        routing_key=routing_key)
