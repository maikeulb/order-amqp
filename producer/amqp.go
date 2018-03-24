package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// var rabbitMQURL = os.Getenv("RABBITMQHOST")
// var amqpURI = "amqp://guest:guest@172.17.0.5:5672/"
var (
	amqpURI = flag.String("amqp", "amqp://guest:guest@172.17.0.5:5672/", "AMQP URI")
)

var conn *amqp.Connection
var ch *amqp.Channel
var q *amqp.Queue

func AddOrderToRabbitMQ(o order) {
	// body := "{'order':" + "'" + orderId.Hex() + "'}"
	payload, err := json.Marshal(o)
	err = ch.Publish(
		"go-amqp-exchange",
		"order-queue",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			// Body:         []byte(body),
			Body:      payload,
			Timestamp: time.Now(),
		})
	log.Printf(" Sent Order %s to queue: %s", o.ID.Hex(), "order_queue")
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func initializeAmqp() {
	flag.Parse()
	var err error

	conn, err = amqp.Dial(*amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	ch.QueueDeclare(
		"order-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare the queue")
}
