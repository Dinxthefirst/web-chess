package game

import "fmt"

func (g *Game) Move(move Move) error {
	piece := g.Board[move.StartSquare]
	if piece.pieceType() == None {
		return fmt.Errorf("no piece at %d", move.StartSquare)
	}
	if piece.color() != g.ColorToMove {
		return fmt.Errorf("wrong color")
	}
	flag, err := g.movePiece(move, piece)
	if err != nil {
		return err
	}

	if g.ColorToMove == White {
		g.ColorToMove = Black
	} else {
		g.ColorToMove = White
		g.fullMoveCounter++
	}

	g.lastMove = Move{move.StartSquare, move.TargetSquare, flag}

	return nil
}

func (g *Game) movePiece(move Move, p Piece) (int, error) {
	moves := g.LegalMoves(move.StartSquare)
	validMove := false
	flag := NoFlag
	for _, m := range moves {
		if m.TargetSquare == move.TargetSquare {
			flag = m.flag
			validMove = true
			break
		}
	}
	if !validMove {
		return flag, fmt.Errorf("invalid move")
	}

	if g.Board[move.TargetSquare].pieceType() == None && p.pieceType() != Pawn {
		g.halfMoveCounter++
	}

	if flag == Castling {
		g.handleCastling(move)
	}

	if flag == EnPassantCapture {
		g.handleEnPassantCapture(move)
	}

	g.Board[move.TargetSquare] = p
	g.Board[move.StartSquare] = Piece{None}

	return flag, nil
}

func (g *Game) handleCastling(move Move) {
	if move.TargetSquare == 6 {
		g.Board[5] = g.Board[7]
		g.Board[7] = Piece{None}
	}
	if move.TargetSquare == 2 {
		g.Board[3] = g.Board[0]
		g.Board[0] = Piece{None}
	}
	if move.TargetSquare == 62 {
		g.Board[61] = g.Board[63]
		g.Board[63] = Piece{None}
	}
	if move.TargetSquare == 58 {
		g.Board[59] = g.Board[56]
		g.Board[56] = Piece{None}
	}
}

func (g *Game) handleEnPassantCapture(move Move) {
	if g.ColorToMove == White {
		g.Board[move.TargetSquare-8] = Piece{None}
	} else {
		g.Board[move.TargetSquare+8] = Piece{None}
	}
}
