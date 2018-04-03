# OrderMQ

(Simple) Distributed application with RabbitMQ and MongoDB. There are two parts
to this application: a publisher (exposed as an API) and a consumer (console
application). The publisher receives orders from the API client, persists them to
MongoDB, and publishes them onto the exchange (as serialized JSON). The consumer
listens for incoming messages from the queue via the routing key and logs the
order to the console.

Technology
----------
* Go
* MongoDB
* RabbitMQ

Endpoints
---------

| Method     | URI                                  | Action                                      |
|------------|--------------------------------------|---------------------------------------------|
| `POST`     | `/api/orders`                        | `Make an order`                             |


Sample Usage
---------------

`http post http://localhost:5000/api/orders product=ipad`
```
{
    "id": "5ac3e5a39039e1051da55d1b", 
    "product": "ipad"
}
...
```
logged to console from the publisher:  
`2018/04/03 16:35:47 Sent order 5ac3e5a39039e1051da55d1b to queue: order_queue`

logged to console from the consumer:  
`2018/04/03 16:35:47 Received a message: {"id":"5ac3e5a39039e1051da55d1b","product":"ipad"}`

Run
---
First go go the publisher and then go to `amqp.go` and point the AMQP uri
to your client; then go to `db.go` and point the Mongo uri to
your server. 

After that has been taken care of,
```
go build (consumer)
./consumer
go build (publisher)
./publisher
Go to http://localhost:5000 and visit the above endpoint.
```

TODO
---
Dockerfile
