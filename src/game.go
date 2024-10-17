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
	g.LoadPositionFromFen(fen)
	return g
}

func (g *Game) LoadPositionFromFen(fen string) error {
	pieces, color, halfMoveCounter, fullMoveCounter, err := parseFen(fen)
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

	g.halfMoveCounter = halfMoveCounter
	g.fullMoveCounter = fullMoveCounter

	g.currentFen = fen
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

func (g *Game) Move(move Move) error {
	piece := g.Board[move.StartSquare]
	if piece.pieceType() == None {
		return fmt.Errorf("no piece at %d", move.StartSquare)
	}
	if piece.color() != g.ColorToMove {
		return fmt.Errorf("wrong color")
	}
	flag, err := g.movePiece(move, piece)
	if err != nil {
		return err
	}

	if g.ColorToMove == White {
		g.ColorToMove = Black
	} else {
		g.ColorToMove = White
		g.fullMoveCounter++
	}

	lastMove := Move{move.StartSquare, move.TargetSquare, flag}
	g.currentFen = g.generateFenFromPosition(lastMove)

	return nil
}

func (g *Game) movePiece(move Move, p Piece) (int, error) {
	moves := g.LegalMoves(move.StartSquare)
	validMove := false
	flag := NoFlag
	for _, m := range moves {
		if m.TargetSquare == move.TargetSquare {
			flag = m.flag
			validMove = true
			break
		}
	}
	if !validMove {
		return flag, fmt.Errorf("invalid move")
	}

	if g.Board[move.TargetSquare].pieceType() == None && p.pieceType() != Pawn {
		g.halfMoveCounter++
	}

	if flag == Castling {
		g.handleCastling(move)
	}

	if flag == EnPassantCapture {
		g.handleEnPassantCapture(move)
	}

	g.Board[move.TargetSquare] = p
	g.Board[move.StartSquare] = Piece{None}

	return flag, nil
}

func (g *Game) handleCastling(move Move) {
	if move.TargetSquare == 6 {
		g.Board[5] = g.Board[7]
		g.Board[7] = Piece{None}
	}
	if move.TargetSquare == 2 {
		g.Board[3] = g.Board[0]
		g.Board[0] = Piece{None}
	}
	if move.TargetSquare == 62 {
		g.Board[61] = g.Board[63]
		g.Board[63] = Piece{None}
	}
	if move.TargetSquare == 58 {
		g.Board[59] = g.Board[56]
		g.Board[56] = Piece{None}
	}
}

func (g *Game) handleEnPassantCapture(move Move) {
	if g.ColorToMove == White {
		g.Board[move.TargetSquare-8] = Piece{None}
	} else {
		g.Board[move.TargetSquare+8] = Piece{None}
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

	splitFen := strings.Split(g.currentFen, " ")
	castlingRights := splitFen[2]
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
