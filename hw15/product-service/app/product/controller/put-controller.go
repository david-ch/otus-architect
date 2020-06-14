package controller

import (
	"encoding/json"
	"github.com/david-ch/otus-architect/hw15/product-service/product/db"
	"github.com/david-ch/otus-architect/hw15/product-service/product/model"
	"net/http"
	"strconv"

	"github.com/david-ch/otus-architect/hw15/product-service/util"
	"github.com/gorilla/mux"
)

type putRq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func OnPut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	rq := &putRq{}
	err = json.NewDecoder(r.Body).Decode(rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	p := &model.Product{
		ID:          id,
		Name:        rq.Name,
		Description: rq.Description,
	}

	err = db.Products.Update(p)
	if err != nil {
		util.Resp(w, http.StatusInternalServerError, util.FromError(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
