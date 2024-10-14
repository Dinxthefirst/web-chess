package game

import "fmt"

type King struct {
	color Color
}

func (k *King) Move(fromRow, fromCol, toRow, toCol int, g *Game) error {
	if checkKing(fromRow, fromCol, toRow, toCol, g) {
		g.Board[fromRow][fromCol] = nil
		g.Board[toRow][toCol] = k
		return nil
	}
	return fmt.Errorf("invalid move for king")
}

func (k *King) Type() string {
	return "King"
}

func (k *King) Color() Color {
	return k.color
}

func (k *King) Symbol() string {
	return PieceImages[k.Type()][k.Color()]
}
