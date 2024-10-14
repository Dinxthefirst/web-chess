package api

import (
	"encoding/json"
	"net/http"

	game "web-chess/src"
)

type GameHandler struct {
	game *game.Game
}

type MoveRequest struct {
	FromX int        `json:"fromX"`
	FromY int        `json:"fromY"`
	ToX   int        `json:"toX"`
	ToY   int        `json:"toY"`
	Color game.Color `json:"color"`
}

func (h *GameHandler) NewGame(w http.ResponseWriter, req *http.Request) {
	h.game = game.NewGame()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.game.ToGameState())
}

func (h *GameHandler) Move(w http.ResponseWriter, req *http.Request) {
	var move MoveRequest

	err := json.NewDecoder(req.Body).Decode(&move)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.game.Move(move.FromX, move.FromY, move.ToX, move.ToY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.game.ToGameState())
}

func (h *GameHandler) CurrentState(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.game.ToGameState())
}
