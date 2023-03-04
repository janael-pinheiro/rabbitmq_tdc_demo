from time import sleep

from pika import BasicProperties
from pika.delivery_mode import DeliveryMode

from rabbitmq.amqp.publisher import AMQPPublisher
from rabbitmq.amqp_demo import (create_channel, create_connection,
                                create_exchange, create_queue)
from rabbitmq.utils.logger import logger_factory


def main() -> None:
    SECONDS_BETWEEN_PUBLICATION: int = 1
    connection = create_connection()
    channel = create_channel(connection)
    exchange_name: str = "demo"
    routing_key: str = "123A"
    logger = logger_factory()
    create_exchange(exchange_name, "fanout", channel)
    create_queue("python", channel, exchange_name, routing_key)

    properties = BasicProperties(delivery_mode=DeliveryMode.Persistent)
    publisher = AMQPPublisher(
        channel=channel,
        exchange_name=exchange_name,
        routing_key=routing_key,
        properties=properties,
        logger=logger)

    while True:
        publisher.publish(content='{"content": "Hello world!"}')
        sleep(SECONDS_BETWEEN_PUBLICATION)

if __name__ == "__main__":
    main()
