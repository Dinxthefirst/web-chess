package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"web-chess/api"
	"web-chess/perft"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: -perft <depth> | -server")
	}

	switch os.Args[1] {
	case "-perft":
		if len(os.Args) < 3 {
			fmt.Println("Usage: -perft <depth>")
			return
		}
		depth, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid depth: %s\n", os.Args[2])
			return
		}
		perft.RunPerft(depth)
	case "-server":
		srv := api.NewServer()
		fmt.Println("SERVER CREATED")
		log.Fatal(http.ListenAndServe("127.0.0.1:42069", srv))
	default:
		fmt.Println("Usage: -perft <depth> | -server")
	}
}
