package event

import (
	"log"
	"math/rand"
)

func onOrderCreated(msg map[string]interface{}) {
	order := msg["order"].(map[string]interface{})
	orderId := order["ID"].(string)

	log.Printf("order created with id %v", orderId)

	var err error
	if rand.Intn(100) >= 0 {
		err = NotifyOrderPaid(orderId)
	} else {
		err = NotifyOrderPaymentFailed(orderId)
	}

	if err != nil {
		log.Printf("Error while processing event 'orderCreated' %v", err)
	}
}
