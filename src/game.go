package game

import (
	"fmt"
	"strconv"
	"strings"
)

const initialFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func NewGame() *Game {
	return NewGameFromFen(initialFen)
}

func NewGameFromFen(fen string) *Game {
	precomputedMoveData()
	g := &Game{}
	g.Board = [BoardSize * BoardSize]Piece{}
	g.loadPositionFromFen(fen)
	return g
}

func (g *Game) loadPositionFromFen(fen string) error {
	pieces, color, castlingRights, halfMoveCounter, fullMoveCounter, err := parseFen(fen)
	if err != nil {
		return err
	}

	g.LoadPiecesFromFen(pieces)

	if color == "w" {
		g.ColorToMove = White
	} else if color == "b" {
		g.ColorToMove = Black
	} else {
		return fmt.Errorf("invalid color to move")
	}

	g.castlingRights = castlingRights

	g.halfMoveCounter = halfMoveCounter
	g.fullMoveCounter = fullMoveCounter

	return nil
}

func (g *Game) LoadPiecesFromFen(fen string) {
	rank := 7
	file := 0
	for _, char := range fen {
		if char == '/' {
			file = 0
			rank--
			continue
		}
		if char >= '1' && char <= '8' {
			file += int(char - '0')
			continue
		}
		g.Board[rank*BoardSize+file] = createPiece(char)
		file++
	}
}

func (g *Game) generateFenFromPosition(lastMove Move) string {
	fen := ""

	for rank := 7; rank >= 0; rank-- {
		emptySquares := 0

		for file := 0; file < 8; file++ {
			i := rank*BoardSize + file
			piece := g.Board[i]
			if piece.pieceType() == None {
				emptySquares++
				continue
			}
			if emptySquares > 0 {
				fen += fmt.Sprintf("%d", emptySquares)
				emptySquares = 0
			}
			fen += symbolForPiece(piece)
		}

		if emptySquares > 0 {
			fen += fmt.Sprintf("%d", emptySquares)
		}

		fen += "/"
	}
	fen = strings.TrimSuffix(fen, "/")

	fen += " "

	if g.ColorToMove == White {
		fen += "w"
	} else {
		fen += "b"
	}
	fen += " "

	castlingRights := g.castlingRights
	if lastMove.flag == Castling {
		castlingRights = updateCastlingRights(castlingRights, lastMove)
	}
	fen += castlingRights

	fen += " "

	if lastMove.flag == PawnTwoForward {
		fen += toChessNotation(enPassantSquare(lastMove, g.ColorToMove))
	} else {
		fen += "-"
	}

	fen += " "

	halfMove := strconv.Itoa(g.halfMoveCounter)
	fen += halfMove

	fen += " "

	fullMove := strconv.Itoa(g.fullMoveCounter)
	fen += fullMove

	return fen
}
