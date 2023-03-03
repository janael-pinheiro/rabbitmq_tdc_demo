from rabbitmq.amqp_demo import (
    create_connection,
    create_channel,
    create_exchange,
    create_queue)
from rabbitmq.amqp.subscriber import Subscriber
from rabbitmq.utils.logger import logger_factory


def main() -> None:
    connection = create_connection()
    channel = create_channel(connection)
    exchange_name: str = "demo"
    routing_key: str = "123A"
    queue_name: str = "python"
    create_exchange(exchange_name, "fanout", channel)
    create_queue(queue_name, channel, exchange_name, routing_key)
    logger = logger_factory()

    subscriber = Subscriber(
        channel=channel,
        queue_name=queue_name,
        routing_key=routing_key,
        logger=logger)

    subscriber.subscribe()

if __name__ == "__main__":
    main()
