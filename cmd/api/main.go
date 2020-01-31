package main

import (
	"log"
	"net/http"

	"github.com/gusibi/nCov/api"
)

func main() {
	log.Println("start server..")
	router := api.GetRouters()

	log.Fatal(http.ListenAndServe(":8080", router))
}
