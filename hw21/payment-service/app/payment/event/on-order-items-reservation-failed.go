package event

import "log"

func onOrderItemsReservationFailed(msg map[string]interface{}) {
	orderId := int64(msg["orderId"].(float64))
	log.Printf("onOrderItemsReservationFailed: payment cancelled for order (id %v)", orderId)

	err := NotifyOrderPaymentFailed(orderId)
	if err != nil {
		log.Printf("Error while processing event 'orderItemsReservationFailed' %v", err)
	}
}
