package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func onIndex(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello from [%s]", hostname)
}

type health struct {
	Status string
}

func onHealth(w http.ResponseWriter, r *http.Request) {
	h := health{"OK"}

	js, _ := json.Marshal(h)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", onIndex)
	http.HandleFunc("/health", onHealth)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
