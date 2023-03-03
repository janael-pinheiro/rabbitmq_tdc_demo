# RabbitMQ demo

## RabbitMQ Docker
```
$ sudo docker pull rabbitmq:3.9.20-management
$ sudo docker run --rm -d -p 5672:5672 -p 15672:15672 rabbitmq:3.9.20-management
```

## Environment variables
- AMQP_URL


```
$ source scripts/set_venv.sh
```

## Python 3
```
$ cd python/
$ poetry install
```

Inside python folder:
```
$ poetry run python3 -m rabbitmq.demo
```