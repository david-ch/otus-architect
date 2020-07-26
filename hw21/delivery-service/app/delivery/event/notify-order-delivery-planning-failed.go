package event

func NotifyOrderDeliveryPlanningFailed(orderId int64) error {
	event := map[string]interface{}{
		"type":  "orderDeliveryPlanningFailed",
		"orderId": orderId,
	}

	return sendEvent(event)
}
