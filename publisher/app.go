package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	Router *mux.Router
	DB     *mgo.Database
}

func (a *App) Initialize() {
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = session.DB("amqp_demo_db")
	a.Router = mux.NewRouter()
	initializeAmqp()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println("Listening on port: 5000")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/orders", a.makeOrder).Methods("POST")
}

func (a *App) makeOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var o order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	o.ID = bson.NewObjectId()
	_, err := AddOrderToMongoDB(a.DB, o)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, o)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
