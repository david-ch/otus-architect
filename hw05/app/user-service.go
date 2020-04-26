package main

import (
	"encoding/json"
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

	r.HandleFunc("/user/{id}", onGet).Methods("GET")
	r.HandleFunc("/user", onPost).Methods("POST")
	r.HandleFunc("/user/{id}", onPut).Methods("PUT")
	r.HandleFunc("/user/{id}", onDelete).Methods("DELETE")
}

func resp(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
