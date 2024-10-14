package game

import "fmt"

type Pawn struct {
	color Color
}

func (p *Pawn) Move(fromX, fromY, toX, toY int, g *Game) error {
	switch p.color {
	case White:
		if fromY == 1 && toY == 3 && fromX == toX {
			if g.Board[fromX][2] == nil && g.Board[fromX][3] == nil {
				g.Board[fromX][fromY] = nil
				g.Board[toX][toY] = p
				return nil
			}
		} else if toY == fromY+1 && fromX == toX {
			if g.Board[fromX][fromY+1] == nil {
				g.Board[fromX][fromY] = nil
				g.Board[toX][toY] = p
				return nil
			}
		}
	case Black:
		if fromY == 7 && toY == 5 && fromX == toX {
			if g.Board[fromX][6] == nil && g.Board[fromX][5] == nil {
				g.Board[fromX][fromY] = nil
				g.Board[toX][toY] = p
				return nil
			}
		} else if toY == fromY-1 && fromX == toX {
			if g.Board[fromX][fromY-1] == nil {
				g.Board[fromX][fromY] = nil
				g.Board[toX][toY] = p
				return nil
			}
		}
	}
	return fmt.Errorf("invalid move for pawn")
}

func (p *Pawn) Type() string {
	return "Pawn"
}

func (p *Pawn) Color() Color {
	return p.color
}

func (p *Pawn) Symbol() string {
	return PieceUnicode[p.Type()][p.Color()]
}
