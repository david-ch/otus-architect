package model

import (
	"time"
)

type Order struct {
	ID                 string
	UserID             int64
	ProductID          int64
	Status             string
	DeliveryTimeslot   time.Time
	CreatedByRequestId string
}
