package game

import "fmt"

type Knight struct {
	color Color
}

func (n *Knight) Move(fromRow, fromCol, toRow, toCol int, g *Game) error {
	if checkKnight(fromRow, fromCol, toRow, toCol, g) {
		g.Board[fromRow][fromCol] = nil
		g.Board[toRow][toCol] = n
		return nil
	}
	return fmt.Errorf("invalid move for knight")
}

func (n *Knight) Type() string {
	return "Knight"
}

func (n *Knight) Color() Color {
	return n.color
}

func (n *Knight) Symbol() string {
	return PieceImages[n.Type()][n.Color()]
}
