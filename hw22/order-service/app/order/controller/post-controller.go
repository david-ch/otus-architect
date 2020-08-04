package controller

import (
	"encoding/json"
	"github.com/david-ch/otus-architect/order-service/order/db"
	"github.com/david-ch/otus-architect/order-service/order/event"
	"github.com/david-ch/otus-architect/order-service/order/model"
	"github.com/david-ch/otus-architect/order-service/util"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type postRq struct {
	UserID           int64     `json:"userId"`
	ProductID        int64     `json:"productId"`
	DeliveryTimeslot time.Time `json:"deliveryTimeslot"`
}

func OnPost(w http.ResponseWriter, r *http.Request) {
	rq := &postRq{}
	err := json.NewDecoder(r.Body).Decode(rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	requestId := r.Header.Get("X-Request-Id")

	o := model.Order{
		ID:                 uuid.New().String(),
		UserID:             rq.UserID,
		ProductID:          rq.ProductID,
		Status:             "pending",
		CreatedByRequestId: requestId,
		DeliveryTimeslot:   rq.DeliveryTimeslot,
	}

	err = db.Orders.Create(o)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	err = event.NotifyOrderCreated(o)
	if err != nil {
		util.Resp(w, http.StatusInternalServerError, util.FromError(err))
		return
	}

	util.Resp(w, http.StatusOK, o.ID)
}
