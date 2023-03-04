from rabbitmq.amqp_demo import (
    create_connection,
    create_channel,
    create_exchange,
    create_queue)
from rabbitmq.amqp.publisher import AMQPPublisher


def main() -> None:
    connection = create_connection()
    channel = create_channel(connection)
    exchange_name: str = "demo"
    routing_key: str = "123A"
    create_exchange(exchange_name, "fanout", channel)
    create_queue("python", channel, exchange_name, routing_key)

    publisher = AMQPPublisher(
        channel=channel,
        exchange_name=exchange_name,
        routing_key=routing_key,
        properties=None)
    
    publisher.publish(content='{"content": "Hello world!"}')

if __name__ == "__main__":
    main()