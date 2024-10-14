package game

import "fmt"

type Queen struct {
	color Color
}

func (q *Queen) Move(fromRow, fromCol, toRow, toCol int, g *Game) error {
	if checkLinear(fromRow, fromCol, toRow, toCol, g) || checkDiagonal(fromRow, fromCol, toRow, toCol, g) {
		g.Board[fromRow][fromCol] = nil
		g.Board[toRow][toCol] = q
		return nil
	}
	return fmt.Errorf("invalid move for queen")
}

func (q *Queen) Type() string {
	return "Queen"
}

func (q *Queen) Color() Color {
	return q.color
}

func (q *Queen) Symbol() string {
	return PieceImages[q.Type()][q.Color()]
}
