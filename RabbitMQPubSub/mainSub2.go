package main

import "rabbitmqdemo/RabbitMQ"

func main() {
	rabbitmq:= RabbitMQ.NewRabbitMQPubSub("penguin")
	rabbitmq.RecieveSub()
}

