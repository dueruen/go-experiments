package sub

import (
	"context"
	"fmt"

	shippingPb "github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/pb"
	"github.com/nats-io/go-nats"
)

const (
	subject = "shipping-events"
	queue   = "shipping-queue"
)

var ordersShipped int32

func GetOrdersShipped() int32 {
	return ordersShipped
}

func Listen(url string, shippingClient shippingPb.ShippingServiceClient) (err error) {
	conn, err := connectToNats(url)
	if err != nil {
		return
	}

	conn.QueueSubscribe(subject, queue, func(e *shippingPb.Event) {
		switch e.EventType {
		case shippingPb.EventType_INVENTORY_RESERVED:
			{
				fmt.Printf("Just got a inventory reserved\n")
				shippingClient.ShipOrder(context.Background(), &shippingPb.ShipOrderRequest{})
			}
		case shippingPb.EventType_ORDER_SHIPPED:
			{
				fmt.Printf("Just got a order shipped\n")
				ordersShipped++
			}
		default:

		}
	})
	return
}

func connectToNats(url string) (encodedConn *nats.EncodedConn, err error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return
	}
	return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
}
