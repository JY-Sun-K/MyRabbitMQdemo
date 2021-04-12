package main

import "rabbitmqdemo/RabbitMQ"

func main() {
	penguinOne := RabbitMQ.NewRabbitMQRouting("exPenguin","penguin_one")
	penguinOne.RecieveRouting()
}
