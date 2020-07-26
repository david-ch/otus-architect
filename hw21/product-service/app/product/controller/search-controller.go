package controller

import (
	"github.com/david-ch/otus-architect/product-service/product/db"
	"github.com/david-ch/otus-architect/product-service/util"
	"github.com/gorilla/mux"
	"net/http"
)

func OnSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	search, ok := vars["search"]
	if !ok {
		return
	}

	productsIds, err := db.Products.StartsWith(search)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	util.Resp(w, http.StatusOK, productsIds)
}