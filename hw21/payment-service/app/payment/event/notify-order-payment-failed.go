package event


func NotifyOrderPaymentFailed(id int64) error {
	event := map[string]interface{}{
		"type":  "orderPaymentFailed",
		"orderId": id,
	}

	return sendEvent(event)
}
