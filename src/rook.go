package game

import "fmt"

type Rook struct {
	color Color
}

func (r *Rook) Move(fromRow, fromCol, toRow, toCol int, g *Game) error {
	if checkLinear(fromRow, fromCol, toRow, toCol, g) {
		g.Board[fromRow][fromCol] = nil
		g.Board[toRow][toCol] = r
		return nil
	}
	return fmt.Errorf("invalid move for rook")
}

func (r *Rook) Type() string {
	return "Rook"
}

func (r *Rook) Color() Color {
	return r.color
}

func (r *Rook) Symbol() string {
	return PieceImages[r.Type()][r.Color()]
}
