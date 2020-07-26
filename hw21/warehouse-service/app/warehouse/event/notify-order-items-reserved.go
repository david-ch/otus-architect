package event

func NotifyOrderItemsReserved(orderId int64) error {
	event := map[string]interface{}{
		"type":  "orderItemsReserved",
		"orderId": orderId,
	}

	return sendEvent(event)
}
