package main

import "gopkg.in/mgo.v2/bson"

type order struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Product string        `bson:"name" json:"product"`
}
