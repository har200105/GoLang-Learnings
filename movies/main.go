package main

import (
	"fmt"
	"log"
	"movies/router"
	"net/http"
)

func main() {

	r := router.Router()

	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Server is Running")

}
