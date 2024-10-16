package test

import (
	"testing"
	game "web-chess/src"
)

func TestNewGameFen(t *testing.T) {
	g := game.NewGame()
	expectedFenString := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	if g.GetFen() != expectedFenString {
		t.Errorf("\nExpected:\n%v\nGot:\n%v", expectedFenString, g.GetFen())
	}
}

func TestGameFenAfterMove(t *testing.T) {
	g := game.NewGame()

	g.Move(game.Move{StartSquare: 12, TargetSquare: 28}) // e4

	expectedFenString := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"

	if g.GetFen() != expectedFenString {
		t.Errorf("\nExpected:\n%v\nGot:\n%v", expectedFenString, g.GetFen())
	}
}
