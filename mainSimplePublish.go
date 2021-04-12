package main

import (
	"fmt"
	"rabbitmqdemo/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("penguin")
	rabbitmq.PublicSimple("hello penguin")
	fmt.Println("发送成功")
}
