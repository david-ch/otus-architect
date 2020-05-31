package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func registerUserRoutes(r *mux.Router) {
	onGet := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			resp(w, http.StatusBadRequest, errorResp{Message: "Cannot parse id"})
			return
		}

		u, err := users.get(id)
		if err != nil {
			resp(w, http.StatusNotFound, fromError(err))
			return
		}

		resp(w, http.StatusOK, u)
	}

	onPost := func(w http.ResponseWriter, r *http.Request) {
		u := &user{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			resp(w, http.StatusBadRequest, errorResp{Message: "Cannot parse request"})
			return
		}

		id, err := users.create(u)
		if err != nil {
			resp(w, http.StatusBadRequest, fromError(err))
			return
		}

		resp(w, http.StatusOK, id)
	}

	onPut := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			resp(w, http.StatusBadRequest, errorResp{Message: "Cannot parse id"})
			return
		}

		u := &user{}
		err = json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			resp(w, http.StatusBadRequest, errorResp{Message: "Cannot parse request"})
			return
		}

		u.ID = id
		err = users.update(u)
		if err != nil {
			resp(w, http.StatusNotFound, fromError(err))
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	onDelete := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			resp(w, http.StatusBadRequest, errorResp{Message: "Cannot parse id"})
			return
		}

		err = users.delete(id)
		if err != nil {
			resp(w, http.StatusNotFound, fromError(err))
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	onCount := func(w http.ResponseWriter, r *http.Request) {
		if rand.Float32() < 0.05 {
			resp(w, http.StatusInternalServerError, errorResp{Message: "Error for test"})
			return
		}

		cnt, err := users.count()

		if err != nil {
			resp(w, http.StatusInternalServerError, errorResp{Message: "Cannot load data"})
			return
		}

		resp(w, http.StatusOK, cnt)
	}

	r.HandleFunc("/user/{id}", withMetrics("/user", onGet)).Methods("GET")
	r.HandleFunc("/user", withMetrics("/user", onPost)).Methods("POST")
	r.HandleFunc("/user/{id}", withMetrics("/user", onPut)).Methods("PUT")
	r.HandleFunc("/user/{id}", withMetrics("/user", onDelete)).Methods("DELETE")
	r.HandleFunc("/userCount", withMetrics("/userCount", onCount)).Methods("GET")
}

func resp(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
