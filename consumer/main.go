package main

import (
	"log"

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

	// Declaring the queue that the consumer will consume from
	queue, err := ch.QueueDeclare(
		"my_queue", // Queue name
		true,       // Durable
		false,      // Should delete when queue is not in use?
		false,      // Exclusive
		false,      // No-wait
		nil,        // No additional arguments
	)
	if err != nil {
		log.Fatal("Error on declaring queue:", err)
	}

	// Consuming queue messages
	messages, err := ch.Consume(
		queue.Name, // Queue name
		"my_app",   // Consumer
		true,       // Auto-ack
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // No additional arguments
	)
	if err != nil {
		log.Fatal("Error registering consumer:", err)
	}

	forever := make(chan bool)

	// Loop to process received messages
	go func() {
		for message := range messages {
			log.Printf("Message received: %s", message.Body)
		}
	}()

	log.Println("[*] Waiting for messages...")
	<-forever
}
