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
	g.ColorToMove = White
	g.halfMoveCounter = 0
	g.fullMoveCounter = 1
	return g
}

func (g *Game) LoadPositionFromFen(fen string) {
	splitFen := strings.Split(fen, " ")
	g.LoadPiecesFromFen(splitFen[0])

	if splitFen[1] == "w" {
		g.ColorToMove = White
	} else {
		g.ColorToMove = Black
	}

	g.currentFen = fen
}

func (g *Game) LoadPiecesFromFen(fen string) {
	rank := 7
	file := 0
	for _, c := range fen {
		if c == '/' {
			file = 0
			rank--
			continue
		}
		if c >= '1' && c <= '8' {
			file += int(c - '0')
			continue
		}
		color := White
		if c >= 'a' && c <= 'z' {
			color = Black
		}
		pieceType := None
		switch c {
		case 'p', 'P':
			pieceType = Pawn
		case 'n', 'N':
			pieceType = Knight
		case 'b', 'B':
			pieceType = Bishop
		case 'r', 'R':
			pieceType = Rook
		case 'q', 'Q':
			pieceType = Queen
		case 'k', 'K':
			pieceType = King
		}
		g.Board[rank*BoardSize+file] = Piece{pieceType | color}
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

	g.generateFenFromPosition(move, flag)

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

	if flag == EnPassantCapture {
		if g.ColorToMove == White {
			g.Board[move.TargetSquare-8] = Piece{None}
		} else {
			g.Board[move.TargetSquare+8] = Piece{None}
		}
	}

	g.Board[move.TargetSquare] = p
	g.Board[move.StartSquare] = Piece{None}
	return flag, nil
}

func (g *Game) LegalMoves(index int) []Move {
	moves := g.generateMoves()

	filteredMoves := []Move{}
	for _, m := range moves {
		if m.StartSquare == index {
			filteredMoves = append(filteredMoves, m)
		}
	}
	return filteredMoves
}

func (g *Game) generateFenFromPosition(lastMove Move, flag int) {
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

	// TODO: castling rights
	splitFen := strings.Split(g.currentFen, " ")
	castlingRights := splitFen[2]
	if flag == Castling {
		if lastMove.TargetSquare == 6 {
			castlingRights = strings.Replace(castlingRights, "K", "", -1)
		}
		if lastMove.TargetSquare == 2 {
			castlingRights = strings.Replace(castlingRights, "Q", "", -1)
		}
		if lastMove.TargetSquare == 58 {
			castlingRights = strings.Replace(castlingRights, "k", "", -1)
		}
		if lastMove.TargetSquare == 62 {
			castlingRights = strings.Replace(castlingRights, "q", "", -1)
		}
		if castlingRights == "" {
			castlingRights = "-"
		}
	}
	fen += castlingRights

	fen += " "

	if flag == PawnTwoForward {
		var enPassantSquare int
		if g.ColorToMove == White {
			enPassantSquare = lastMove.TargetSquare + 8
		} else {
			enPassantSquare = lastMove.TargetSquare - 8
		}
		fen += toChessNotation(enPassantSquare)
	} else {
		fen += "-"
	}

	fen += " "

	halfMove := strconv.Itoa(g.halfMoveCounter)
	fen += halfMove

	fen += " "

	fullMove := strconv.Itoa(g.fullMoveCounter)
	fen += fullMove

	g.currentFen = fen
}
