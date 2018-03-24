package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var username string
var address []string
var session *mgo.Session
var asession *mgo.Session
var collection *mgo.Collection
var serr error

// var mongoURL = os.Getenv("MONGOHOST")
var mongoURL = "172.17.0.4"

// var mongoDB = os.Getenv("MONGODB")
var mongoDB = "amqpdemo"

const (
	COLLECTION = "orders"
)

func AddOrderToMongoDB(db *mgo.Database, o order) (bson.ObjectId, error) {

	err := db.C(COLLECTION).Insert(&o)

	AddOrderToRabbitMQ(o)
	return o.ID, err
}
