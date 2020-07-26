package controller

import (
	"github.com/david-ch/otus-architect/product-service/product/cache"
	"github.com/david-ch/otus-architect/product-service/product/db"
	"github.com/david-ch/otus-architect/product-service/util"
	"net/http"
)

func OnGetPromoted(w http.ResponseWriter, r *http.Request) {
	productsIds, found := cache.GetPromoted()

	if !found {
		productsIds, err := db.Products.GetPromoted()
		if err != nil {
			util.Resp(w, http.StatusNotFound, util.FromError(err))
			return
		}

		cache.CachePromoted(productsIds)
	}

	util.Resp(w, http.StatusOK, productsIds)
}
