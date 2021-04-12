package main

import (
	"fmt"
	"rabbitmqdemo/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq:=RabbitMQ.NewRabbitMQPubSub(""+"penguin")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishPub("订阅模式产生第"+strconv.Itoa(i)+"条数据")
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
}
