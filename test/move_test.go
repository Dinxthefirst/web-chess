package test

import (
	"fmt"
	"testing"
	game "web-chess/src"
)

func compareMovesErrorMessage(expected, got []game.Move) string {
	return fmt.Sprintf("\nExpected:\n%v\nGot:\n%v", expected, got)
}

func movesEqualIgnoreFlag(m1, m2 game.Move) bool {
	return m1.StartSquare == m2.StartSquare && m1.TargetSquare == m2.TargetSquare
}

func movesEqual(first, second []game.Move) bool {
	if len(first) != len(second) {
		return false
	}

	for _, firstMove := range first {
		found := false
		for _, secondMove := range second {
			if movesEqualIgnoreFlag(firstMove, secondMove) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func TestKingMoves(t *testing.T) {
	kingFen := "4k3/8/8/8/8/8/P7/4K3 w - - 0 1"

	g := game.NewGameFromFen(kingFen)

	kingMoves := g.LegalMovesAtIndex(4)

	expectedMoves := []game.Move{
		{StartSquare: 4, TargetSquare: 3},
		{StartSquare: 4, TargetSquare: 5},
		{StartSquare: 4, TargetSquare: 11},
		{StartSquare: 4, TargetSquare: 12},
		{StartSquare: 4, TargetSquare: 13},
	}

	if !movesEqual(expectedMoves, kingMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, kingMoves))
	}
}

func TestKingMovesFromCorner(t *testing.T) {
	kingFen := "4k3/8/8/8/8/8/8/K6R w - - 0 1"

	g := game.NewGameFromFen(kingFen)

	kingMoves := g.LegalMovesAtIndex(0)

	expectedMoves := []game.Move{
		{StartSquare: 0, TargetSquare: 1},
		{StartSquare: 0, TargetSquare: 8},
		{StartSquare: 0, TargetSquare: 9},
	}

	if !movesEqual(expectedMoves, kingMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, kingMoves))
	}
}

func TestRookMoves(t *testing.T) {
	rookFen := "4k3/8/8/8/P7/8/8/R3K3 w - - 0 1"

	g := game.NewGameFromFen(rookFen)

	rookMoves := g.LegalMovesAtIndex(0)

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
	knightFen := "4k3/8/8/8/4N3/8/8/4K3 w - - 0 1"

	g := game.NewGameFromFen(knightFen)

	knightMoves := g.LegalMovesAtIndex(28)

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
		t.Error(compareMovesErrorMessage(expectedMoves, knightMoves))
	}
}

func TestKnightMovesStart(t *testing.T) {
	g := game.NewGame()

	knightMoves := g.LegalMovesAtIndex(1)

	expectedMoves := []game.Move{
		{StartSquare: 1, TargetSquare: 16},
		{StartSquare: 1, TargetSquare: 18},
	}

	if !movesEqual(expectedMoves, knightMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, knightMoves))
	}
}

func TestPawnMovesStart(t *testing.T) {
	g := game.NewGame()

	pawnMoves := g.LegalMovesAtIndex(8)

	expectedMoves := []game.Move{
		{StartSquare: 8, TargetSquare: 16},
		{StartSquare: 8, TargetSquare: 24},
	}

	if !movesEqual(expectedMoves, pawnMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, pawnMoves))
	}
}

func TestPawnMovesAfterStart(t *testing.T) {
	g := game.NewGame()

	g.Move(game.Move{StartSquare: 8, TargetSquare: 16})
	g.Move(game.Move{StartSquare: 51, TargetSquare: 43})

	pawnMoves := g.LegalMovesAtIndex(16)

	expectedMoves := []game.Move{
		{StartSquare: 16, TargetSquare: 24},
	}

	if !movesEqual(expectedMoves, pawnMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, pawnMoves))
	}
}

func TestPawnDiagonalCapture(t *testing.T) {
	pawnFen := "4k3/8/8/8/5p2/4P3/8/4K3 w - - 0 1"

	g := game.NewGameFromFen(pawnFen)

	pawnMoves := g.LegalMovesAtIndex(20)

	expectedMoves := []game.Move{
		{StartSquare: 20, TargetSquare: 28},
		{StartSquare: 20, TargetSquare: 29},
	}

	if !movesEqual(expectedMoves, pawnMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, pawnMoves))
	}
}

func TestPawnDiagonalCaptureAgain(t *testing.T) {
	pawnFen := "3k4/8/8/8/r7/6r1/7P/3K4 w - - 0 1"

	g := game.NewGameFromFen(pawnFen)

	pawnMoves := g.LegalMovesAtIndex(15)

	expectedMoves := []game.Move{
		{StartSquare: 15, TargetSquare: 22},
		{StartSquare: 15, TargetSquare: 23},
		{StartSquare: 15, TargetSquare: 31},
	}

	if !movesEqual(expectedMoves, pawnMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, pawnMoves))
	}
}

func TestBishopMoves(t *testing.T) {
	bishopFen := "4k3/8/8/8/8/8/1B7/4K3 w - - 0 1"

	g := game.NewGameFromFen(bishopFen)

	bishopMoves := g.LegalMovesAtIndex(9)

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
		t.Error(compareMovesErrorMessage(expectedMoves, bishopMoves))
	}
}

func TestQueenMoves(t *testing.T) {
	queenFen := "4k3/8/8/8/8/8/8/1Q2K3 w - - 0 1"

	g := game.NewGameFromFen(queenFen)

	queenMoves := g.LegalMovesAtIndex(1)

	expectedMoves := []game.Move{
		{StartSquare: 1, TargetSquare: 0},
		{StartSquare: 1, TargetSquare: 2},
		{StartSquare: 1, TargetSquare: 3},
		{StartSquare: 1, TargetSquare: 8},
		{StartSquare: 1, TargetSquare: 9},
		{StartSquare: 1, TargetSquare: 17},
		{StartSquare: 1, TargetSquare: 25},
		{StartSquare: 1, TargetSquare: 33},
		{StartSquare: 1, TargetSquare: 41},
		{StartSquare: 1, TargetSquare: 49},
		{StartSquare: 1, TargetSquare: 57},
		{StartSquare: 1, TargetSquare: 10},
		{StartSquare: 1, TargetSquare: 19},
		{StartSquare: 1, TargetSquare: 28},
		{StartSquare: 1, TargetSquare: 37},
		{StartSquare: 1, TargetSquare: 46},
		{StartSquare: 1, TargetSquare: 55},
	}

	if !movesEqual(expectedMoves, queenMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, queenMoves))
	}
}

func TestPawnEnPassant(t *testing.T) {
	g := game.NewGame()

	g.Move(game.Move{StartSquare: 12, TargetSquare: 28}) // e4
	g.Move(game.Move{StartSquare: 52, TargetSquare: 44}) // e6
	g.Move(game.Move{StartSquare: 28, TargetSquare: 36}) // e5
	g.Move(game.Move{StartSquare: 51, TargetSquare: 35}) // d5

	pawnMoves := g.LegalMovesAtIndex(36)

	expectedMoves := []game.Move{
		{StartSquare: 36, TargetSquare: 43},
	}

	if !movesEqual(expectedMoves, pawnMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, pawnMoves))
	}
}

func TestCastling(t *testing.T) {
	castlingFen := "4k3/8/8/8/8/8/8/R3K3 w Q - 0 1"

	g := game.NewGameFromFen(castlingFen)

	kingMoves := g.LegalMovesAtIndex(4)

	expectedMoves := []game.Move{
		{StartSquare: 4, TargetSquare: 3},
		{StartSquare: 4, TargetSquare: 5},
		{StartSquare: 4, TargetSquare: 11},
		{StartSquare: 4, TargetSquare: 12},
		{StartSquare: 4, TargetSquare: 13},
		{StartSquare: 4, TargetSquare: 2},
	}

	if !movesEqual(expectedMoves, kingMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, kingMoves))
	}
}

func TestChecking(t *testing.T) {
	checkingFen := "7k/R7/8/8/8/8/8/K7 w - - 0 1"

	g := game.NewGameFromFen(checkingFen)

	g.Move(game.Move{StartSquare: 48, TargetSquare: 56}) // Ra8

	kingMoves := g.LegalMovesAtIndex(63)

	expectedMoves := []game.Move{
		{StartSquare: 63, TargetSquare: 54},
		{StartSquare: 63, TargetSquare: 55},
	}

	if !movesEqual(expectedMoves, kingMoves) {
		t.Error(compareMovesErrorMessage(expectedMoves, kingMoves))
	}
}
