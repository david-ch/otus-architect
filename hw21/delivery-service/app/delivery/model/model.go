package model

import "time"

type DeliveryPlan struct {
	OrderId  int64
	Timeslot time.Time
}
