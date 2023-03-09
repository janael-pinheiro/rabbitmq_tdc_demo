# RabbitMQ demo

This repository contains code for demo asynchronous communication through RabbitMQ. Therefore, this code is highly experimental. Contributions are more than welcome.

## Demonstrated features
- Automatic reconnection;
- Publication confirmation;
- Data preservation;
- Manual acknowledgment;
- Setting the prefetch count.

## RabbitMQ Docker
```
$ sudo docker pull rabbitmq:3.9.20-management
$ sudo docker run --rm -d -p 5672:5672 -p 15672:15672 rabbitmq:3.9.20-management
```
You can access the RabbitMQ web interface in your browser on port 15672 (username and password: guest).

## Environment variables
- AMQP_URL;
- PREFETCH_COUNT;
- HOST;
- VHOST;
- PORT;
- USERNAME;
- PASSWORD;
- EXCHANGE;
- QUEUE;
- ROUNTING_KEY.

```
$ source scripts/set_venv.sh
```

## Python 3
Inside the root folder of the project:
```
$ cd python/
$ poetry install
```

```
$ poetry run python3 -m rabbitmq.subscriber_demo
$ poetry run python3 -m rabbitmq.publisher_demo
```

## Go
Inside the root folder of the project:
```
$ cd go/
$ go mod tidy
```
```
go run subscriber/main.go
go run publisher/main.go
```