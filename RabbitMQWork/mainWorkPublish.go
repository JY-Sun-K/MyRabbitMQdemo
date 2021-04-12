package main

import (
	"fmt"
	"rabbitmqdemo/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq:=RabbitMQ.NewRabbitMQSimple(""+"penguin")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublicSimple("hello penguin!"+strconv.Itoa(i))
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
}