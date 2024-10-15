package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	game "web-chess/src"
)

type GameHandler struct {
	game *game.Game
}

func (h *GameHandler) NewGame(w http.ResponseWriter, req *http.Request) {
	h.game = game.NewGame()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.game)
}

func (h *GameHandler) Move(w http.ResponseWriter, req *http.Request) {
	var move game.Move

	err := json.NewDecoder(req.Body).Decode(&move)
	if err != nil {
		fmt.Printf("Error decoding move: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.game.Move(move)
	if err != nil {
		fmt.Printf("Error moving piece: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.game)
}

func (h *GameHandler) CurrentState(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.game)
}
