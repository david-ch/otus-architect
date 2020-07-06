package controller

import (
	"github.com/david-ch/otus-architect/hw15/product-service/product/cache"
	"github.com/david-ch/otus-architect/hw15/product-service/product/db"
	"log"
	"net/http"
	"strconv"

	"github.com/david-ch/otus-architect/hw15/product-service/util"
	"github.com/gorilla/mux"
)

type getRs struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Promoted    bool   `json:"promoted"`
}

func OnGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		util.Resp(w, http.StatusBadRequest, util.FromError(err))
		return
	}

	p, found := cache.GetProduct(id)
	if !found {
		p, err = db.Products.Get(id)
		if err != nil {
			util.Resp(w, http.StatusNotFound, util.FromError(err))
			return
		}
		cache.CacheProduct(p)
	} else {
		log.Print("DEBUG Found in cache ", id)
	}

	rs := getRs{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Promoted:    p.Promoted,
	}

	util.Resp(w, http.StatusOK, rs)
}

