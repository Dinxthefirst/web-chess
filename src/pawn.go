package game

import "fmt"

type Pawn struct {
	color Color
}

func (p *Pawn) Move(fromRow, fromCol, toRow, toCol int, g *Game) error {
	switch p.color {
	case White:
		if fromRow == 6 && toRow == 4 && fromCol == toCol {
			if g.Board[5][fromCol] == nil && g.Board[4][toCol] == nil {
				g.Board[fromRow][fromCol] = nil
				g.Board[toRow][toCol] = p
				return nil
			}
		} else if toRow == fromRow-1 && fromCol == toCol {
			if g.Board[toRow][toCol] == nil {
				g.Board[fromRow][fromCol] = nil
				g.Board[toRow][toCol] = p
				return nil
			}
		}
	case Black:
		if fromRow == 1 && toRow == 3 && fromCol == toCol {
			if g.Board[2][fromCol] == nil && g.Board[3][fromCol] == nil {
				g.Board[fromRow][fromCol] = nil
				g.Board[toRow][toCol] = p
				return nil
			}
		} else if toRow == fromRow+1 && fromCol == toCol {
			if g.Board[toRow][toCol] == nil {
				g.Board[fromRow][fromCol] = nil
				g.Board[toRow][toCol] = p
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
	return PieceImages[p.Type()][p.Color()]
}
