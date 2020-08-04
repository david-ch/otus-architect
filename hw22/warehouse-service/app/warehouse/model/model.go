package model

type Stock struct {
	ProductId int64
	Amount    int64
}

type StockReservation struct {
	ProductId int64
	OrderId   string
	Amount    int64
}
