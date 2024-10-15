package game

import (
	"fmt"
)

// const initialFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const initialFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

func NewGame() *Game {
	precomputedMoveData()
	g := &Game{}
	g.Board = [BoardSize * BoardSize]Piece{}
	g.LoadPositionFromFen(initialFen)
	g.ColorToMove = White
	return g
}

func (g *Game) LoadPositionFromFen(fen string) {
	rank := 7
	file := 0
	for _, c := range fen {
		if c == '/' {
			file = 0
			rank--
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
}

func (g *Game) Move(move Move) error {
	piece := g.Board[move.StartSquare]
	if piece.pieceType() == None {
		return fmt.Errorf("no piece at %d", move.StartSquare)
	}
	if piece.color() != g.ColorToMove {
		return fmt.Errorf("wrong color")
	}
	err := g.movePiece(move, piece)
	if err != nil {
		return err
	}

	if g.ColorToMove == White {
		g.ColorToMove = Black
	} else {
		g.ColorToMove = White
	}
	return nil
}

func (g *Game) movePiece(move Move, p Piece) error {
	moves := g.LegalMoves(move.StartSquare)
	validMove := false
	for _, m := range moves {
		if m.TargetSquare == move.TargetSquare {
			validMove = true
			break
		}
	}
	if !validMove {
		return fmt.Errorf("invalid move")
	}

	g.Board[move.TargetSquare] = p
	g.Board[move.StartSquare] = Piece{None}
	return nil
}

func (g *Game) LegalMoves(index int) []Move {
	moves := g.generateMoves()

	filteredMoves := []Move{}
	for _, m := range moves {
		if m.StartSquare == index {
			filteredMoves = append(filteredMoves, m)
		}
	}
	return filteredMoves
}
