package event

import (
	"github.com/david-ch/otus-architect/delivery-service/client"
	"github.com/david-ch/otus-architect/delivery-service/delivery/db"
	"github.com/david-ch/otus-architect/delivery-service/delivery/model"
	"log"
)

func onOrderItemsReserved(msg map[string]interface{}) {
	orderId := msg["orderId"].(string)

	order, err := client.GetOrder(orderId)
	if err != nil {
		log.Printf("onOrderPaid: error while loading order with id %v. %v", orderId, err)
		NotifyOrderDeliveryPlanningFailed(orderId)
		return
	}

	deliveryPlan := model.DeliveryPlan{
		OrderId: orderId,
		Timeslot: order.DeliveryTimeslot,
	}
	db.Delivery.CreateDeliveryPlan(deliveryPlan)
	NotifyOrderDeliveryPlanned(orderId)
	log.Printf("onOrderItemsReserved: timeslot reserved for order (id %v)", orderId)
}
