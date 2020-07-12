package controller

import (
	"encoding/json"
	"github.com/david-ch/otus-architect/hw15/product-service/product/db"
	"github.com/david-ch/otus-architect/hw15/product-service/product/model"
	"net/http"

	"github.com/david-ch/otus-architect/hw15/product-service/util"
)

type postRq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func OnPost(w http.ResponseWriter, r *http.Request) {
	rq := &postRq{}
	err := json.NewDecoder(r.Body).Decode(rq)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	p := &model.Product{
		Name:        rq.Name,
		Description: rq.Description,
	}

	id, err := db.Products.Create(p)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	util.Resp(w, http.StatusOK, id)
}