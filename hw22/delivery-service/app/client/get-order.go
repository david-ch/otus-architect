package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Order struct {
	ID               string    `json:"id"`
	UserID           int64     `json:"userId"`
	ProductID        int64     `json:"productId"`
	Status           string    `json:"status"`
	DeliveryTimeslot time.Time `json:"deliveryTimeslot"`
}

// GetOrder loads order information from order-service
func GetOrder(orderId string) (order Order, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/%v", os.Getenv("ORDER_SVC_ADDRESS"), orderId))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &order)
	if err != nil {
		return
	}

	return
}
