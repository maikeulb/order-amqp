# OrderMQ

Distributed messaging demo application with RabbitMQ and MongoDB. Registered
deliveries (through an API) are stored in MongoDB and published onto an
exchange. The order messages (serialized JSON) are then pushed onto a queue (from the
exchange) where they wait to be consumed by a dummy service (console
application).

Protocol: AMQP 0-9-1 (Advanced Messaging Queue Protocol)
Exchange Type: Direct (routed)

Technology
----------
* Go
* MongoDb
* RabbitMQ

Endpoints
---------

| Method     | URI                                  | Action                                      |
|------------|--------------------------------------|---------------------------------------------|
| `POST`     | `/api/orders`                        | `Make an order`                             |


Sample Usage
---------------

`http post http://localhost:5000/api/order product=book`
```
    "networks": [
        {
            "id": "bbbike", 
            "location": {
                "city": "Bielsko-Biała", 
                "country": "PL", 
                "latitude": 49.8225, 
                "longitude": 19.044444
            }, 
            "name": "BBBike"
        }, 
        {
            "id": "bixi-montreal", 
            "location": {
                "city": "Montreal, QC", 
                "country": "CA", 
                "latitude": 45.5086699, 
                "longitude": -73.55399249999999
            }, 
            "name": "Bixi"
        }, 
...
```
logged to console from the publisher: `retrieved networks from remote api in 455.69015ms`

logged to console from the subscriber: `retrieved networks from cache in 6.075998ms`

Run
---

First go go the publisher and then go to `amqp.go` and point the AMQP uri
variable to your client; then go to `db.go` and point the Mongo variables to
your server. 

After that has been taken care of,
```
go build (consumer)
./consumer
go build (producer)
./producer
Go to http://localhost:5000 and visit the above endpoint.
```

TODO
---
Dockerfile
