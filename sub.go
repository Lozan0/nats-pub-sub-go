package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func subscribe(ready chan bool) {

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	sub, _ := nc.SubscribeSync("msg.test")

	// Avisamos que la suscripción está lista.
	ready <- true

	m, err := sub.NextMsg(3 * time.Second)
	if err == nil {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	} else {
		fmt.Println("NextMsg timed out.")
	}
}
