# RabbitMQ demo

## Demonstrated features
- Automatic reconnection;
- Publication confirmation;
- Data preservation.

## RabbitMQ Docker
```
$ sudo docker pull rabbitmq:3.9.20-management
$ sudo docker run --rm -d -p 5672:5672 -p 15672:15672 rabbitmq:3.9.20-management
```

## Environment variables
- AMQP_URL.

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