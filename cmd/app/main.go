package main

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

type TimeResponse struct {
	UTC string `json:"utc"`
	OS  string `json:"os"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC()
	osType := runtime.GOOS
	response := TimeResponse{
		UTC: currentTime.Format(time.RFC3339),
		OS:  osType,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", timeHandler)

	port := ":8088"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
