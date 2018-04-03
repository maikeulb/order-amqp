package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var (
	amqpURI = flag.String("amqp", "amqp://guest:guest@172.17.0.4:5672/", "AMQP URI")
)

var conn *amqp.Connection
var ch *amqp.Channel
var replies <-chan amqp.Delivery

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func run() {
	flag.Parse()
	var err error
	var q amqp.Queue

	conn, err = amqp.Dial(*amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Printf("got Connection, getting Channel...")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer conn.Close()

	q, err = ch.QueueDeclare(
		"order-queue", // name
		true,          // durable
		false,         // delete when usused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Error declaring the Queue")

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		q.Name, q.Messages, q.Consumers, "go-test-key")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Error consuming the Queue")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
