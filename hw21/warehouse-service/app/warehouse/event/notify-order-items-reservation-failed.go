package event

func NotifyOrderItemsReservationFailed(orderId int64) error {
	event := map[string]interface{}{
		"type":  "orderItemsReservationFailed",
		"orderId": orderId,
	}

	return sendEvent(event)
}
