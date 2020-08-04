package controller

import (
	"github.com/david-ch/otus-architect/order-service/order/db"
	"github.com/david-ch/otus-architect/order-service/util"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type getRs struct {
	ID               string    `json:"id"`
	UserID           int64     `json:"userId"`
	ProductID        int64     `json:"productId"`
	Status           string    `json:"status"`
	DeliveryTimeslot time.Time `json:"deliveryTimeslot"`
}

func OnGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		util.Resp(w, http.StatusBadRequest, "GET request should have an id")
		return
	}

	o, err := db.Orders.Get(id)
	if err != nil {
		util.Resp(w, http.StatusNotFound, util.FromError(err))
		return
	}

	rs := getRs{
		ID:               o.ID,
		UserID:           o.UserID,
		ProductID:        o.ProductID,
		Status:           o.Status,
		DeliveryTimeslot: o.DeliveryTimeslot,
	}

	util.Resp(w, http.StatusOK, rs)
}
