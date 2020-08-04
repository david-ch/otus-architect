package model

import "time"

type DeliveryPlan struct {
	OrderId  string
	Timeslot time.Time
}
