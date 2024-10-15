package game

import "fmt"

const initialFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

func NewGame() *Game {
	g := GameFromFen(initialFen)
	return g
}

func GameFromFen(fen string) *Game {
	g := &Game{}
	g.Board = [BoardSize * BoardSize]Piece{}
	g.ActiveColor = White
	rank := 0
	file := 0
	for _, c := range fen {
		if c == '/' {
			rank++
			file = 0
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
	return g
}

func (g *Game) Move(move Move) error {
	piece := g.Board[move.FromRow*BoardSize+move.FromCol]
	if piece.pieceType() == None {
		return fmt.Errorf("no piece at %d %d", move.FromRow, move.FromCol)
	}
	if piece.color() != g.ActiveColor {
		return fmt.Errorf("wrong color")
	}

	err := g.movePiece(move, piece)
	if err != nil {
		return err
	}

	g.ActiveColor = 1 - g.ActiveColor
	return nil
}

func (g *Game) movePiece(move Move, p Piece) error {
	return fmt.Errorf("not implemented")
	// switch p.Type & 7 {
	// case Pawn:
	// 	return g.movePawn(move)
	// case Knight:
	// 	return g.moveKnight(move)
	// case Bishop:
	// 	return g.moveBishop(move)
	// case Rook:
	// 	return g.moveRook(move)
	// case Queen:
	// 	return g.moveQueen(move)
	// case King:
	// 	return g.moveKing(move)
	// case None:
	// 	return fmt.Errorf("no piece")
	// }
	// return fmt.Errorf("invalid piece type")
}
