package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GetOrder loads order information from order-service
func GetOrder(orderId int64) (order map[string]interface{}, err error) {
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
