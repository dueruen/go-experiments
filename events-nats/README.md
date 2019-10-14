# Events and NATS experiment
A experiment to see how NATS can be used and also using the auto-generated go types from grpc as the object model.

## Inspired by
https://www.youtube.com/watch?v=DzGuDNHsOQ0&list=PLUn4JBwKdy9tEjKW1pNIyw6klSXwz1bgk&index=3&t=1857s

## Run the experiment
Start NATS:
```bash
docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
```
Start the services:
```bash
go run order_service/cmd/order-server/main.go
go run shipping_service/cmd/shipping-server/main.go
```
I have not made a client so a cli client can bu used, like [Evans](https://github.com/ktr0731/evans)
```bash
$ evans order_service/pkg/transport/pb/chat.proto
$ call CreateOrder
{}
$ call GetShippedOrders
{
  "amount": 1
}

```
The GetshippedOrders returns the amount of shipped orders this count is increased when the order service receives a event that the order is shipped.
## Auto-Generate using the proto file
```bach
cd order_service/pkg/transport/pb
protoc order.proto --go_out=plugins=grpc:.
cd shipping_service/pkg/transport/pb
protoc shipping.proto --go_out=plugins=grpc:.
```
## Stop the experiment
```bash
docker container stop nats-main
```