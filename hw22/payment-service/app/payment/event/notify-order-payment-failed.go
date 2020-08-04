package event

func NotifyOrderPaymentFailed(id string) error {
	event := map[string]interface{}{
		"type":    "orderPaymentFailed",
		"orderId": id,
	}

	return sendEvent(event)
}
