package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"fmt"

	"html/template"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	a "pack.ag/amqp"
)

var amqpclient *a.Client
var aerr error
var amqpsession *a.Session
var amqpsender *a.Sender

// var hosts string
// var db string

var rabbitMQURL = os.Getenv("RABBITMQHOST")

func AddOrder(order Order) (orderId string) {

	return orderId
}

func AddOrderToMongoDB(order Order) (orderId string) {

	session = asession.Copy()

	NewOrderID := bson.NewObjectId()

	order.ID = NewOrderID.Hex()

	order.Status = "Open"
	if order.Source == "" || order.Source == "string" {
		order.Source = os.Getenv("SOURCE")
	}

	database = "amqp_test2"
	password = "" //V2

	log.Print(mongoURL, "AMQP")

	defer session.Close()

	serr = collection.Insert(order)
	log.Println("_id:", order)

	if serr != nil {
		log.Fatal("Problem inserting data: ", serr)
		log.Println("_id:", order)
	}

	return order.ID
}

func AddOrderToRabbitMQ(orderId string, orderSource string) {

	conn, err := amqp.Dial(rabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"order", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "{'order':" + "'" + orderId + "', 'source':" + "'" + orderSource + "'}"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		})
	log.Printf(" [x] Sent %s " + body + " queue:" + q.Name)
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func InitializeRMQ() {

	if len(os.Getenv("RABBITMQHOST")) == 0 {
		log.Print("The environment variable RABBITMQHOST has not been set")
	} else {
		log.Print("The environment variable RABBITMQHOST is " + os.Getenv("RABBITMQHOST"))
	}
}
