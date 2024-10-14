package game

import (
	"fmt"
)

func NewGame() *Game {
	g := &Game{
		ActiveColor: White,
	}
	g.initialPosition()
	return g
}

func (g *Game) initialPosition() {
	for i := 0; i < BoardSize; i++ {
		g.Board[1][i] = &Pawn{color: Black}
		g.Board[6][i] = &Pawn{color: White}
	}
	// g.Board[0][0] = &Rook{color: White}
	// g.Board[0][7] = &Rook{color: White}
	// g.Board[7][0] = &Rook{color: Black}
	// g.Board[7][7] = &Rook{color: Black}
}

func (g *Game) Move(fromX, fromY, toX, toY int) error {
	piece := g.Board[fromX][fromY]
	if piece == nil {
		return fmt.Errorf("no piece at %d %d", fromX, fromY)
	}
	if piece.Color() != g.ActiveColor {
		return fmt.Errorf("wrong color")
	}

	err := piece.Move(fromX, fromY, toX, toY, g)
	if err != nil {
		return err
	}
	g.ActiveColor = 1 - g.ActiveColor
	return nil
}
