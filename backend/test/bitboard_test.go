package test

import (
	"fmt"
	"testing"
	game "web-chess/backend/src"
)

func bitboardString(bitboard uint64) string {
	str := ""
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			str += fmt.Sprintf("%b", bitboard&(1<<(rank*8+file))>>(rank*8+file))
		}
		str += "\n"
	}
	return str
}

func bitboardsEqual(bitboards1, bitboards2 [23]uint64) ([]int, bool) {
	equal := true
	indices := []int{}
	for i := 0; i < 8; i++ {
		if bitboards1[i] != bitboards2[i] {
			equal = false
			indices = append(indices, i)
		}
	}
	return indices, equal
}

func handleBitBoardMismatch(bitboards [23]uint64, bitboardsAfter [23]uint64, indices []int, t *testing.T) {
	errorString := ""
	for i := range indices {
		bitboardStr := bitboardString(bitboards[i])
		bitboardAfterStr := bitboardString(bitboardsAfter[i])
		errorString += fmt.Sprintf("\nBitboard[%d] not equal:\n%s!=\n%s", i, bitboardStr, bitboardAfterStr)
	}
	t.Error(errorString)
}

func TestBitBoardMakingMove(t *testing.T) {
	g := game.NewGame()

	move := game.Move{StartSquare: 8, TargetSquare: 16}
	g.Move(move)

	bitboards := g.BitBoards()

	g = game.NewGameFromFen("rnbqkbnr/pppppppp/8/8/8/P7/1PPPPPPP/RNBQKBNR w KQkq - 0 1")

	bitboardsAfter := g.BitBoards()

	indices, equal := bitboardsEqual(bitboards, bitboardsAfter)
	if !equal {
		handleBitBoardMismatch(bitboards, bitboardsAfter, indices, t)
	}

}

func TestBitBoardAfterOneMove(t *testing.T) {
	g := game.NewGame()

	bitboards := g.BitBoards()

	move := game.Move{StartSquare: 8, TargetSquare: 16}
	g.Move(move)
	g.UnmakeMove(move)

	bitboardsAfter := g.BitBoards()

	indices, equal := bitboardsEqual(bitboards, bitboardsAfter)
	if !equal {
		handleBitBoardMismatch(bitboards, bitboardsAfter, indices, t)
	}
}

func TestBitBoardAfterCastling(t *testing.T) {
	castleFen := "r3k2r/8/8/8/8/8/8/R3K2R w KQkq - 0 1"
	g := game.NewGameFromFen(castleFen)

	bitboards := g.BitBoards()

	move := game.Move{StartSquare: 4, TargetSquare: 6, Flag: game.Castling}
	g.Move(move)
	g.UnmakeMove(move)

	bitboardsAfter := g.BitBoards()

	indices, equal := bitboardsEqual(bitboards, bitboardsAfter)
	if !equal {
		handleBitBoardMismatch(bitboards, bitboardsAfter, indices, t)
	}
}

func TestBitBoardAfterCapture(t *testing.T) {
	captureFen := "4k3/8/8/8/8/8/4p3/4K3 w - - 0 1"
	g := game.NewGameFromFen(captureFen)

	bitboards := g.BitBoards()

	move := game.Move{StartSquare: 4, TargetSquare: 12}
	g.Move(move)
	g.UnmakeMove(move)

	bitboardsAfter := g.BitBoards()

	indices, equal := bitboardsEqual(bitboards, bitboardsAfter)
	if !equal {
		handleBitBoardMismatch(bitboards, bitboardsAfter, indices, t)
	}
}

func TestBitBoardAfterEnPassantCapture(t *testing.T) {
	enPassantFen := "4k3/8/8/3Pp3/8/8/8/4Q3 w - e6 0 1"
	g := game.NewGameFromFen(enPassantFen)

	bitboards := g.BitBoards()

	move := game.Move{StartSquare: 35, TargetSquare: 44}
	g.Move(move)
	g.UnmakeMove(move)

	bitboardsAfter := g.BitBoards()

	indices, equal := bitboardsEqual(bitboards, bitboardsAfter)
	if !equal {
		handleBitBoardMismatch(bitboards, bitboardsAfter, indices, t)
	}
}

func TestBitBoardAfterPromotion(t *testing.T) {
	promotionFen := "4k3/P7/8/8/8/8/8/4K3 w - - 0 1"

	g := game.NewGameFromFen(promotionFen)

	bitboards := g.BitBoards()

	move := game.Move{StartSquare: 48, TargetSquare: 56, Flag: game.PromoteToQueen}
	g.Move(move)
	g.UnmakeMove(move)

	bitboardsAfter := g.BitBoards()

	indices, equal := bitboardsEqual(bitboards, bitboardsAfter)
	if !equal {
		handleBitBoardMismatch(bitboards, bitboardsAfter, indices, t)
	}
}

func TestBitBoardAfterPawnTwoForward(t *testing.T) {
	pawnTwoForwardFen := "4k3/8/8/8/8/8/P7/4K3 w - - 0 1"

	g := game.NewGameFromFen(pawnTwoForwardFen)

	bitboards := g.BitBoards()

	move := game.Move{StartSquare: 8, TargetSquare: 16}
	g.Move(move)
	g.UnmakeMove(move)

	bitboardsAfter := g.BitBoards()

	indices, equal := bitboardsEqual(bitboards, bitboardsAfter)
	if !equal {
		handleBitBoardMismatch(bitboards, bitboardsAfter, indices, t)
	}
}
