package event

func NotifyOrderDeliveryPlanned(orderId string) error {
	event := map[string]interface{}{
		"type":    "orderDeliveryPlanned",
		"orderId": orderId,
	}

	return sendEvent(event)
}
