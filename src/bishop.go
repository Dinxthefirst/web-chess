package game

import "fmt"

type Bishop struct {
	color Color
}

func (b *Bishop) Move(fromRow, fromCol, toRow, toCol int, g *Game) error {
	if checkDiagonal(fromRow, fromCol, toRow, toCol, g) {
		g.Board[fromRow][fromCol] = nil
		g.Board[toRow][toCol] = b
		return nil
	}
	return fmt.Errorf("invalid move for bishop")
}

func (b *Bishop) Type() string {
	return "Bishop"
}

func (b *Bishop) Color() Color {
	return b.color
}

func (b *Bishop) Symbol() string {
	return PieceImages[b.Type()][b.Color()]
}
