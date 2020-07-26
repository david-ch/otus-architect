package event


func NotifyOrderPaid(id int64) error {
	event := map[string]interface{}{
		"type":  "orderPaid",
		"orderId": id,
	}

	return sendEvent(event)
}
