package main

import (
	"testing"
	game "web-chess/src"
)

func movesEqual(first, second []game.Move) bool {
	if len(first) != len(second) {
		return false
	}

	exists := make(map[game.Move]bool)
	for _, value := range first {
		exists[value] = true
	}
	for _, value := range second {
		if !exists[value] {
			return false
		}
	}
	return true
}

func TestKingMoves(t *testing.T) {
	kingFen := "4k3/8/8/8/8/8/P7/4K3"

	g := game.NewGameFromFen(kingFen)

	kingMoves := g.LegalMoves(4)
	t.Log("king moves:", kingMoves)

	expectedMoves := []game.Move{
		{StartSquare: 4, TargetSquare: 3},
		{StartSquare: 4, TargetSquare: 5},
		{StartSquare: 4, TargetSquare: 11},
		{StartSquare: 4, TargetSquare: 12},
		{StartSquare: 4, TargetSquare: 13},
	}

	if !movesEqual(expectedMoves, kingMoves) {
		t.Errorf("Expected %v, got %v", expectedMoves, kingMoves)
	}
}

func TestRookMoves(t *testing.T) {
	rookFen := "4k3/8/8/8/P7/8/8/R2rK3"

	g := game.NewGameFromFen(rookFen)

	rookMoves := g.LegalMoves(0)
	t.Log("rook moves:", rookMoves)

	expectedMoves := []game.Move{
		{StartSquare: 0, TargetSquare: 1},
		{StartSquare: 0, TargetSquare: 2},
		{StartSquare: 0, TargetSquare: 3},
		{StartSquare: 0, TargetSquare: 8},
		{StartSquare: 0, TargetSquare: 16},
	}

	if !movesEqual(expectedMoves, rookMoves) {
		t.Errorf("Expected %v, got %v", expectedMoves, rookMoves)
	}
}

func TestKnightMoves(t *testing.T) {
	knightFen := "4k3/8/8/8/4N3/8/8/4K3"

	g := game.NewGameFromFen(knightFen)

	knightMoves := g.LegalMoves(28)
	t.Log("knight moves:", knightMoves)

	expectedMoves := []game.Move{
		{StartSquare: 28, TargetSquare: 43},
		{StartSquare: 28, TargetSquare: 45},
		{StartSquare: 28, TargetSquare: 34},
		{StartSquare: 28, TargetSquare: 38},
		{StartSquare: 28, TargetSquare: 18},
		{StartSquare: 28, TargetSquare: 22},
		{StartSquare: 28, TargetSquare: 11},
		{StartSquare: 28, TargetSquare: 13},
	}

	if !movesEqual(expectedMoves, knightMoves) {
		t.Errorf("Expected %v, got %v", expectedMoves, knightMoves)
	}
}

func TestKnightMovesStart(t *testing.T) {
	g := game.NewGame()

	knightMoves := g.LegalMoves(1)
	t.Log("knight moves:", knightMoves)

	expectedMoves := []game.Move{
		{StartSquare: 1, TargetSquare: 16},
		{StartSquare: 1, TargetSquare: 18},
	}

	if !movesEqual(expectedMoves, knightMoves) {
		t.Errorf("Expected %v, got %v", expectedMoves, knightMoves)
	}
}

func TestPawnMovesStart(t *testing.T) {
	g := game.NewGame()

	pawnMoves := g.LegalMoves(8)
	t.Log("pawn moves:", pawnMoves)

	expectedMoves := []game.Move{
		{StartSquare: 8, TargetSquare: 16},
		{StartSquare: 8, TargetSquare: 24},
	}

	if !movesEqual(expectedMoves, pawnMoves) {
		t.Errorf("Expected %v, got %v", expectedMoves, pawnMoves)
	}
}

func TestPawnMovesAfterStart(t *testing.T) {
	g := game.NewGame()

	g.Move(game.Move{StartSquare: 8, TargetSquare: 16})
	g.Move(game.Move{StartSquare: 51, TargetSquare: 43})

	pawnMoves := g.LegalMoves(16)
	t.Log("pawn moves:", pawnMoves)

	expectedMoves := []game.Move{
		{StartSquare: 16, TargetSquare: 24},
	}

	if !movesEqual(expectedMoves, pawnMoves) {
		t.Errorf("Expected %v, got %v", expectedMoves, pawnMoves)
	}
}

func TestPawnDiagonalCapture(t *testing.T) {
	pawnFen := "4k3/8/8/8/5p2/4P3/8/4K3"

	g := game.NewGameFromFen(pawnFen)

	pawnMoves := g.LegalMoves(20)
	t.Log("pawn moves:", pawnMoves)

	expectedMoves := []game.Move{
		{StartSquare: 20, TargetSquare: 28},
		{StartSquare: 20, TargetSquare: 29},
	}

	if !movesEqual(expectedMoves, pawnMoves) {
		t.Errorf("Expected %v, got %v", expectedMoves, pawnMoves)
	}
}

func TestBishopMoves(t *testing.T) {
	bishopFen := "4k3/8/8/8/8/8/1B7/4K3"

	g := game.NewGameFromFen(bishopFen)

	bishopMoves := g.LegalMoves(9)
	t.Log("bishop moves:", bishopMoves)

	expectedMoves := []game.Move{
		{StartSquare: 9, TargetSquare: 0},
		{StartSquare: 9, TargetSquare: 2},
		{StartSquare: 9, TargetSquare: 16},
		{StartSquare: 9, TargetSquare: 18},
		{StartSquare: 9, TargetSquare: 27},
		{StartSquare: 9, TargetSquare: 36},
		{StartSquare: 9, TargetSquare: 45},
		{StartSquare: 9, TargetSquare: 54},
		{StartSquare: 9, TargetSquare: 63},
	}

	if !movesEqual(expectedMoves, bishopMoves) {
		t.Errorf("Expected %v, got %v", expectedMoves, bishopMoves)
	}
}
