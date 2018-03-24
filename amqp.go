package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"gopkg.in/mgo.v2/bson"
	a "pack.ag/amqp"
)

var amqpclient *a.Client
var aerr error
var amqpsession *a.Session
var amqpsender *a.Sender

// var hosts string
// var db string

// var rabbitMQURL = os.Getenv("RABBITMQHOST")
// var rabbitMQURL = flag.String("amqp", "amqp://guest:guest@172.17.0.5:5672/", "AMQP URI")
var rabbitMQURL = "amqp://guest:guest@172.17.0.5:5672/"

func AddOrder(o order) (orderId string) {

	return orderId
}

func AddOrderToMongoDB(o order) (orderId bson.ObjectId) {

	// session = asession.Copy()

	// NewOrderID := bson.NewObjectId()

	// o.ID = NewOrderID.Hex()

	// // o.Status = "Open"
	// if o.Source == "" || o.Source == "string" {
	//  o.Source = os.Getenv("SOURCE")
	// }

	// database = "amqp_demo"
	// password = ""

	// defer session.Close()

	// serr = collection.Insert(o)
	// log.Println("_id:", o)

	// if serr != nil {
	// log.Fatal("Problem inserting data: ", serr)
	// log.Println("_id:", o)
	// }

	// return o.ID
	AddOrderToRabbitMQ(o.ID, "test")
	return o.ID
}

func AddOrderToRabbitMQ(orderId bson.ObjectId, orderSource string) {

	conn, err := amqp.Dial(rabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"order",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	fmt.Println(orderId)
	fmt.Println(orderId.Hex())
	body := "{'order':" + "'" + orderId.Hex() + "', 'source':" + "'" + orderSource + "'}"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		})
	log.Printf(" [x] Sent %s  queue: %s", orderId.Hex(), q.Name)
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
