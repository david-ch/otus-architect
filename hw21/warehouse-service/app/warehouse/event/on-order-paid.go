package event

import (
	"github.com/david-ch/otus-architect/warehouse-service/client"
	"github.com/david-ch/otus-architect/warehouse-service/warehouse/db"
	"github.com/david-ch/otus-architect/warehouse-service/warehouse/model"
	"log"
)

func onOrderPaid(msg map[string]interface{}) {
	orderId := int64(msg["orderId"].(float64))

	order, err := client.GetOrder(orderId)
	if err != nil {
		log.Printf("onOrderPaid: error while loading order with id %v. %v", orderId, err)
		NotifyOrderItemsReservationFailed(orderId)
		return
	}

	productId := int64(order["productId"].(float64))
	stock, err := db.Warehouse.GetStock(productId)
	if err != nil {
		log.Printf("onOrderPaid: error while loading stock for product id %v. %v", productId, err)
		NotifyOrderItemsReservationFailed(orderId)
		return
	}

	if stock.Amount == 0 {
		log.Printf("onOrderPaid: error while reserving product(id %v) for order(id %v). Out of stock",
			productId,
			orderId)
		NotifyOrderItemsReservationFailed(orderId)
		return
	}

	stock.Amount = stock.Amount - 1
	err = db.Warehouse.UpdateStock(stock)
	if err != nil {
		log.Printf("onOrderPaid: error while updating stock for product id %v. %v", productId, err)
		NotifyOrderItemsReservationFailed(orderId)
		return
	}

	reservation := model.StockReservation{
		ProductId: productId,
		OrderId: orderId,
		Amount: 1,
	}
	db.Warehouse.CreateReservation(reservation)

	NotifyOrderItemsReserved(orderId)

	log.Printf("onOrderPaid: product(id %v) successfully reserved for order (id %v)", productId, orderId)
}
