# OrderMQ

(Simple) Distributed application with RabbitMQ and MongoDB. Registered orders
(through an API) are stored in MongoDB and published onto an exchange as
serialized JSON and pushed onto the queue. A dummy service (console application)
listens for income messages from the queue and consumes them.

Protocol: AMQP 0-9-1 (Advanced Messaging Queue Protocol)
Exchange Type: Direct (routed)

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
logged to console from the publisher: `2018/04/03 16:35:47 Sent order 5ac3e5a39039e1051da55d1b to queue: order_queue`
logged to console from the consumer: `2018/04/03 16:35:47 Received a message: {"id":"5ac3e5a39039e1051da55d1b","product":"ipad"}`

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
./producer
Go to http://localhost:5000 and visit the above endpoint.
```

TODO
---
Dockerfile
