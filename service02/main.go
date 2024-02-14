package main

import (
	"log"
	"trab02/rabbitMQ"
)

func main() {
	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	go rabbitMQ.ReceiveAndGenerateToken()
	rabbitMQ.ReceiveAndValidateToken()
}
