package event

func NotifyOrderItemsReserved(orderId string) error {
	event := map[string]interface{}{
		"type":    "orderItemsReserved",
		"orderId": orderId,
	}

	return sendEvent(event)
}
