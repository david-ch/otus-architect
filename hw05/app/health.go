package main

import (
	"encoding/json"
	"net/http"
)

type health struct {
	Status string
}

func onHealth(w http.ResponseWriter, r *http.Request) {
	h := health{"OK"}

	js, _ := json.Marshal(h)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
