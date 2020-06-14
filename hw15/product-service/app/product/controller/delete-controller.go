package controller

import (
	"github.com/david-ch/otus-architect/hw15/product-service/product/db"
	"net/http"
	"strconv"

	"github.com/david-ch/otus-architect/hw15/product-service/util"
	"github.com/gorilla/mux"
)

func OnDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	err = db.Products.Delete(id)
	if err != nil {
		util.Resp(w, http.StatusNotFound, util.FromError(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
