package event

func NotifyOrderPaid(id string) error {
	event := map[string]interface{}{
		"type":    "orderPaid",
		"orderId": id,
	}

	return sendEvent(event)
}
