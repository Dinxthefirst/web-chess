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
		fmt.Println("Usage: -perft <position> <depth> | -server")
	}

	switch os.Args[1] {
	case "-perft", "---p":
		if len(os.Args) < 4 {
			fmt.Println("Usage: -perft <position> <depth>")
			return
		}
		position, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid position: %s\n", os.Args[2])
			return
		}
		depth, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("Invalid depth: %s\n", os.Args[2])
			return
		}
		perft.RunPerft(position, depth)
	case "-server", "---s":
		srv := api.NewServer()
		fmt.Println("SERVER CREATED")
		log.Fatal(http.ListenAndServe("127.0.0.1:42069", srv))
	default:
		fmt.Println("Usage: -perft <position> <depth> | -server")
	}
}
