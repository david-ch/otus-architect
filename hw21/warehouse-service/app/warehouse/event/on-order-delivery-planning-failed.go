package event

import (
	"github.com/david-ch/otus-architect/warehouse-service/warehouse/db"
	"log"
)

func onOrderDeliveryPlanningFailed(msg map[string]interface{}) {
	orderId := int64(msg["orderId"].(float64))

	reservation, err := db.Warehouse.GetReservation(orderId)
	if err != nil {
		log.Printf("onOrderDeliveryPlanningFailed: error while loading reservation for order with id %v. %v", orderId, err)
		return
	}

	stock, err := db.Warehouse.GetStock(reservation.ProductId)
	if err != nil {
		log.Printf("onOrderDeliveryPlanningFailed: error while loading stock for product with id %v. %v", reservation.ProductId, err)
		return
	}

	stock.Amount = stock.Amount + reservation.Amount
	err = db.Warehouse.UpdateStock(stock)
	if err != nil {
		log.Printf("onOrderDeliveryPlanningFailed: error while updating stock for product with id %v. %v", reservation.ProductId, err)
		return
	}

	err = db.Warehouse.DeleteReservation(orderId)
	if err != nil {
		log.Printf("onOrderDeliveryPlanningFailed: error while deleting reservation for order with id %v. %v", orderId, err)
		return
	}

	NotifyOrderItemsReservationFailed(orderId)
}
