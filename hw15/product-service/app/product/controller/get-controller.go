package controller

import (
	"github.com/david-ch/otus-architect/hw15/product-service/product/db"
	"net/http"
	"strconv"

	"github.com/david-ch/otus-architect/hw15/product-service/util"
	"github.com/gorilla/mux"
)

type getRs struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func OnGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	p, err := db.Products.Get(id)
	if err != nil {
		util.Resp(w, http.StatusNotFound, util.FromError(err))
		return
	}

	rs := getRs{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}

	util.Resp(w, http.StatusOK, rs)
}
