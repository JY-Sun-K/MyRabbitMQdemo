package main

import "rabbitmqdemo/RabbitMQ"

func main() {
	penguinTwo := RabbitMQ.NewRabbitMQRouting("exPenguin","penguin_two")
	penguinTwo.RecieveRouting()
}
