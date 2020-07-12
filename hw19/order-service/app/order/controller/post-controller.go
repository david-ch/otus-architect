package controller

import (
	"encoding/json"
	"github.com/david-ch/otus-architect/hw19/order-service/order/db"
	"github.com/david-ch/otus-architect/hw19/order-service/order/model"
	"github.com/david-ch/otus-architect/hw19/order-service/util"
	"net/http"
)

type postRq struct {
	UserID    int64 `json:"userId"`
	ProductID int64 `json:"productId"`
}

func OnPost(w http.ResponseWriter, r *http.Request) {
	rq := &postRq{}
	err := json.NewDecoder(r.Body).Decode(rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	requestId := r.Header.Get("X-Request-Id")

	p := &model.Order{
		UserID:             rq.UserID,
		ProductID:          rq.ProductID,
		Status:             "initialized",
		CreatedByRequestId: requestId,
	}

	id, err := db.Orders.Create(p)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	util.Resp(w, http.StatusOK, id)
}
