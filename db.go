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

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

var mongoURL = os.Getenv("MONGOHOST")

func InitializeMDB() {

	if len(os.Getenv("MONGOHOST")) == 0 {
		log.Print("The environment variable MONGOHOST has not been set")
	} else {
		log.Print("The environment variable MONGOHOST is " + os.Getenv("MONGOHOST"))
	}

	if strings.Contains(mongoURL, "?ssl=true") {

		url, err := url.Parse(mongoURL)
		if err != nil {
			log.Fatal("Problem parsing url: ", err)
		}

		log.Print("user ", url.User)
		// DialInfo holds options for establishing a session with a MongoDB cluster.
		st := fmt.Sprintf("%s", url.User)
		co := strings.Index(st, ":")

		database = st[:co]
		password = st[co+1:]
		log.Print("db ", database, " pwd ", password)
	}

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s.documents.azure.com:10255", database)}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: database, // It can be anything
		Username: database, // Username
		Password: password, // PASSWORD
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}

	session.SetSafe(nil)

	collection = session.DB(database).C("orders")
}
