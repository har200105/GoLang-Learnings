package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endPointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Success",
		Body:   "Hello akipiD",
	}
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.Handle("/api/token-bucket", tokenBucketRateLimiter(endPointHandler))
	http.Handle("/api/per-client", perClientRateLimitter(endPointHandler))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Println("Something Went Wrong.")
	}
}
