package event

import (
	"github.com/david-ch/otus-architect/order-service/order/model"
)

func NotifyOrderCreated(o model.Order) error {
	event := map[string]interface{}{
		"type":  "orderCreated",
		"order": o,
	}

	return sendEvent(event)
}