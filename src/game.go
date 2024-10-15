package game

import "fmt"

func NewGame() *Game {
	g := &Game{
		ActiveColor: White,
	}
	g.initialPosition()
	return g
}

func (g *Game) initialPosition() {
	for i := 0; i < BoardSize; i++ {
		g.Board[6][i] = Piece{Type: Pawn | White}
		g.Board[1][i] = Piece{Type: Pawn | Black}
	}
	// g.Board[7][0] = &Rook{color: White}
	// g.Board[7][7] = &Rook{color: White}
	// g.Board[0][0] = &Rook{color: Black}
	// g.Board[0][7] = &Rook{color: Black}

	// g.Board[7][1] = &Knight{color: White}
	// g.Board[7][6] = &Knight{color: White}
	// g.Board[0][1] = &Knight{color: Black}
	// g.Board[0][6] = &Knight{color: Black}

	// g.Board[7][2] = &Bishop{color: White}
	// g.Board[7][5] = &Bishop{color: White}
	// g.Board[0][2] = &Bishop{color: Black}
	// g.Board[0][5] = &Bishop{color: Black}

	// g.Board[7][3] = &Queen{color: White}
	// g.Board[0][3] = &Queen{color: Black}

	// g.Board[7][4] = &King{color: White}
	// g.Board[0][4] = &King{color: Black}
}

func (g *Game) Move(move Move) error {
	piece := g.Board[move.FromRow][move.FromCol]
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
