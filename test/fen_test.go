package test

import (
	"testing"
	game "web-chess/src"
)

func compareFenStringErrorMessage(expected, got string) string {
	return "Expected:\n" + expected + "\nGot:\n" + got
}

func TestNewGameFen(t *testing.T) {
	g := game.NewGame()
	expectedFenString := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	if g.CurrentFen() != expectedFenString {
		t.Error(compareFenStringErrorMessage(expectedFenString, g.CurrentFen()))
	}
}

func TestFenAfterMove(t *testing.T) {
	g := game.NewGame()

	g.Move(game.Move{StartSquare: 12, TargetSquare: 28}) // e4

	expectedFenString := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"

	if g.CurrentFen() != expectedFenString {
		t.Error(compareFenStringErrorMessage(expectedFenString, g.CurrentFen()))
	}
}

func TestFenAfterMultipleMoves(t *testing.T) {
	g := game.NewGame()

	g.Move(game.Move{StartSquare: 12, TargetSquare: 28}) // e4
	g.Move(game.Move{StartSquare: 50, TargetSquare: 34}) // c5

	expectedFenString := "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"

	if g.CurrentFen() != expectedFenString {
		t.Errorf("\nExpected:\n%v\nGot:\n%v", expectedFenString, g.CurrentFen())
	}

	g.Move(game.Move{StartSquare: 6, TargetSquare: 21}) // Nf3

	expectedFenString = "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"

	if g.CurrentFen() != expectedFenString {
		t.Error(compareFenStringErrorMessage(expectedFenString, g.CurrentFen()))
	}
}

func TestFenAfterCastling(t *testing.T) {
	castlingFen := "4k3/8/8/8/8/8/8/R3K3 w Q - 0 1"
	g := game.NewGameFromFen(castlingFen)

	err := g.Move(game.Move{StartSquare: 4, TargetSquare: 2}) // O-O-O

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expectedFenString := "4k3/8/8/8/8/8/8/2KR4 b - - 1 1"

	if g.CurrentFen() != expectedFenString {
		t.Error(compareFenStringErrorMessage(expectedFenString, g.CurrentFen()))
	}
}

func TestFenAfterEnPassant(t *testing.T) {
	g := game.NewGame()

	g.Move(game.Move{StartSquare: 12, TargetSquare: 28}) // e4
	g.Move(game.Move{StartSquare: 52, TargetSquare: 44}) // e6
	g.Move(game.Move{StartSquare: 28, TargetSquare: 36}) // e5
	g.Move(game.Move{StartSquare: 51, TargetSquare: 35}) // d5
	g.Move(game.Move{StartSquare: 36, TargetSquare: 43}) // exd5

	expectedFenString := "rnbqkbnr/ppp2ppp/3Pp3/8/8/8/PPPP1PPP/RNBQKBNR b KQkq - 0 3"

	if g.CurrentFen() != expectedFenString {
		t.Error(compareFenStringErrorMessage(expectedFenString, g.CurrentFen()))
	}
}

func TestFenAfterUnmakingMove(t *testing.T) {
	initialFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	g := game.NewGameFromFen(initialFen)

	g.Move(game.Move{StartSquare: 12, TargetSquare: 28})       // e4
	g.UnmakeMove(game.Move{StartSquare: 12, TargetSquare: 28}) // e4

	if g.CurrentFen() != initialFen {
		t.Error(compareFenStringErrorMessage(initialFen, g.CurrentFen()))
	}
}
