package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// Connects on RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Error connecting on RabbitMQ:", err)
	}
	defer conn.Close() // defer means: Closes connection in the end of the program

	// Open a communication channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error opening RabbitMQ channel:", err)
	}
	defer ch.Close()

	// Declare an exchange of type "topic"
	// In this case, the exchange is beeing created programatically
	// but it's not required, this exchange can be already created.
	err = ch.ExchangeDeclare(
		"my_exchange", // Exchange name
		"topic",       // Exchange type
		true,          // Durable
		false,         // Auto-delete
		false,         // It's not an internal exchange
		false,         // No-wait
		nil,           // No additional arguments
	)
	if err != nil {
		log.Fatal("Error declaring exchange:", err)
	}

	// Defining messages
	if len(os.Args) < 2 {
		log.Fatal("A message is required, usage: go run publisher/main.go <message> <routing_key(optional)>")
	}
	message := os.Args[1]

	// Defining routing key (optional)
	routingKey := "" // Default
	if len(os.Args) > 2 {
		routingKey = os.Args[2]
	}

	// Publishing the message
	err = ch.Publish(
		"my_exchange", // Exchange name
		routingKey,    // Routing key (optional)
		false,         // Fouding a queue is not mandatory
		false,         // Delivery immediatly is not mandatory
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatal("Error publishing the message, try again later:", err)
	}

	fmt.Printf("Message sent: %s with routing key: %s\n", message, routingKey)
}
