package event

import (
	"github.com/david-ch/otus-architect/order-service/order/db"
	"log"
)

func onOrderPaymentFailed(msg map[string]interface{}) {
	orderId := msg["orderId"].(string)
	log.Printf("order payment failed. order id = %v", orderId)

	order, err := db.Orders.Get(orderId)
	if err != nil {
		log.Printf("onOrderPaymentFailed: error while loading order with id %v. %v", orderId, err)
		return
	}

	order.Status = "cancelled"
	err = db.Orders.Update(order)
	if err != nil {
		log.Printf("onOrderPaymentFailed: error while updating order with id %v. %v", orderId, err)
		return
	}
}
