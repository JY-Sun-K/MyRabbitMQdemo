package main

import (
	"rabbitmqdemo/RabbitMQ"
)

func main() {
	rabbitmq:=RabbitMQ.NewRabbitMQSimple("penguin")
	rabbitmq.ConsumeSimple()
}
