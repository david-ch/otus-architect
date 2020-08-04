package event

func NotifyOrderDeliveryPlanningFailed(orderId string) error {
	event := map[string]interface{}{
		"type":    "orderDeliveryPlanningFailed",
		"orderId": orderId,
	}

	return sendEvent(event)
}
