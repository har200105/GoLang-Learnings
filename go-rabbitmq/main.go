package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func publishMessage(ch *amqp.Channel, queueName, message string) {
	err := ch.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		panic("Failed to publish the message.")
	}
	fmt.Printf(" [x] Sent %s\n", message)
}

func consumeMessages(ch *amqp.Channel, queueName string) {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-acknowledge
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		panic("Failed to register the consumer.")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf(" [x] Received %s\n", d.Body)
		}
	}()

	fmt.Println(" Waiting for messages.")
	<-forever
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}
	defer ch.Close()

	queueName := "akipiD_queue"

	// Declare a queue
	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		panic("Failed to declare a queue")
	}

	// Consume messages from the queue
	go consumeMessages(ch, queueName)

	// Publishing message to the queue
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Hello, akipiD! %d", i)
		publishMessage(ch, queueName, message)
		time.Sleep(time.Second * 2)
	}
	// Allow some time for consuming messages before exiting
	time.Sleep(2 * time.Second)
}
