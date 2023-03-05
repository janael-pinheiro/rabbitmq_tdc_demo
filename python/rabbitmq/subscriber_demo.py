from pika.exchange_type import ExchangeType

from rabbitmq.amqp.subscriber import Subscriber
from rabbitmq.amqp_demo import (create_channel, create_connection,
                                create_exchange, create_queue)
from rabbitmq.utils.logger import logger_factory


def main() -> None:
    connection = create_connection()
    channel = create_channel(connection)
    exchange_name: str = "demo"
    routing_key: str = "123A"
    queue_name: str = "python"
    dead_letter_exchange: str = f"{exchange_name}_dlx"
    queue_dlq_name: str = f"{queue_name}_dlq"
    create_exchange(exchange_name, ExchangeType.fanout, channel)
    create_exchange(dead_letter_exchange, ExchangeType.fanout, channel)
    create_queue(
        queue_name,
        channel,
        exchange_name,
        routing_key,
        dead_letter_exchange)
    create_queue(
        queue_dlq_name,
        channel,
        dead_letter_exchange,
        routing_key
    )
    logger = logger_factory()

    subscriber = Subscriber(
        channel=channel,
        queue_name=queue_name,
        routing_key=routing_key,
        logger=logger)
    subscriber.subscribe()

if __name__ == "__main__":
    main()
