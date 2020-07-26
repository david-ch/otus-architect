package event

func NotifyOrderDeliveryPlanned(orderId int64) error {
	event := map[string]interface{}{
		"type":  "orderDeliveryPlanned",
		"orderId": orderId,
	}

	return sendEvent(event)
}
