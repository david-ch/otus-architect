package event

func NotifyOrderItemsReservationFailed(orderId string) error {
	event := map[string]interface{}{
		"type":    "orderItemsReservationFailed",
		"orderId": orderId,
	}

	return sendEvent(event)
}
