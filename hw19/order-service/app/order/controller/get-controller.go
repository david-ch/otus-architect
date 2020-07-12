package controller

import (
	"github.com/david-ch/otus-architect/hw19/order-service/order/db"
	"github.com/david-ch/otus-architect/hw19/order-service/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type getRs struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userId"`
	ProductID int64  `json:"productId"`
	Status    string `json:"status"`
}

func OnGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	o, err := db.Orders.Get(id)
	if err != nil {
		util.Resp(w, http.StatusNotFound, util.FromError(err))
		return
	}

	rs := getRs{
		ID:        o.ID,
		UserID:    o.UserID,
		ProductID: o.ProductID,
		Status:    o.Status,
	}

	util.Resp(w, http.StatusOK, rs)
}
