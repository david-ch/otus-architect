package model

type Order struct {
	ID                 int64
	UserID             int64
	ProductID          int64
	Status             string
	CreatedByRequestId string
}
