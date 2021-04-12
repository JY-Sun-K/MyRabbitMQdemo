package main

import (
	"fmt"
	"rabbitmqdemo/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	penguinOne := RabbitMQ.NewRabbitMQRouting("exPenguin","penguin_one")
	penguinTwo := RabbitMQ.NewRabbitMQRouting("exPenguin","penguin_two")
	for i := 0; i <=10; i++ {
		penguinOne.PublishRouting("hello pOne"+strconv.Itoa(i))
		penguinTwo.PublishRouting("hello pTwo"+strconv.Itoa(i))
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
}
