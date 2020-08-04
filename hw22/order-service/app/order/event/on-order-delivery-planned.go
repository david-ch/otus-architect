package event

import (
	"github.com/david-ch/otus-architect/order-service/order/db"
	"log"
)

func onOrderDeliveryPlanned(msg map[string]interface{}) {
	orderId := msg["orderId"].(string)
	log.Printf("order delivery planned. order id = %v", orderId)

	order, err := db.Orders.Get(orderId)
	if err != nil {
		log.Printf("onOrderDeliveryPlanned: error while loading order with id %v. %v", orderId, err)
		return
	}

	order.Status = "delivering"
	err = db.Orders.Update(order)
	if err != nil {
		log.Printf("onOrderDeliveryPlanned: error while updating order with id %v. %v", orderId, err)
		return
	}
}