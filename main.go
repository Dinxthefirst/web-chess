package main

import (
	"fmt"
	"log"
	"net/http"
	"web-chess/api"
)

func main() {
	srv := api.NewServer()
	fmt.Println("SERVER CREATED")
	log.Fatal(http.ListenAndServe("127.0.0.1:42069", srv))
}
