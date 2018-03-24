package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

var (
	database string
	password string
	status   string
)

var username string
var address []string
var session *mgo.Session
var asession *mgo.Session
var collection *mgo.Collection
var serr error

var hosts string
var db string

// var mongoURL = os.Getenv("MONGOHOST")
var mongoURL = "172.17.0.4"

// var mongoDB = os.Getenv("MONGODB")
var mongoDB = "amqpdemo"

func InitializeMDB() {

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatal(err)
	}
	collection = session.DB(database).C("orders")
}
