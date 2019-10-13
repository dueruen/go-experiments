package sub

import (
	"log"

	pb "github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/pb"
	"github.com/nats-io/go-nats"
)

const (
	subject = "shipping-events"
	queue   = "shipping-queue"
)

type eventSender struct {
	natsConn *nats.EncodedConn
}

func NewEventSender(url string) (*eventSender, error) {
	conn, err := connectToNats(url)
	if err != nil {
		return nil, err
	}
	return &eventSender{conn}, nil
}

func (e *eventSender) SendEventInventoryReserved() {
	err := e.natsConn.Publish(subject, pb.Event{EventType: pb.EventType_INVENTORY_RESERVED})
	if err != nil {
		log.Fatal(err)
	}
}

func (e *eventSender) SendEventOrderShipped() {
	err := e.natsConn.Publish(subject, pb.Event{EventType: pb.EventType_ORDER_SHIPPED})
	if err != nil {
		log.Fatal(err)
	}
}

func connectToNats(url string) (encodedConn *nats.EncodedConn, err error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return
	}
	return nats.NewEncodedConn(conn, nats.JSON_ENCODER)
}
