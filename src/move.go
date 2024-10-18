package game

import "fmt"

func (g *Game) Move(move Move) error {
	if g.Board[move.StartSquare].pieceType() == None {
		return fmt.Errorf("no piece at %d", move.StartSquare)
	}

	moves := g.GenerateMoves()
	validMove := false
	for _, m := range moves {
		if m.StartSquare == move.StartSquare && m.TargetSquare == move.TargetSquare {
			validMove = true
			move = m
			break
		}
	}
	if !validMove {
		return fmt.Errorf("no move from %d to %d", move.StartSquare, move.TargetSquare)
	}

	return g.MakeMove(move)
}

func (g *Game) MakeMove(move Move) error {
	var currentGameState uint32 = 0
	originalCastleRights := g.currentGameState & 15
	newCastleState := originalCastleRights

	moveFrom := move.StartSquare
	moveTo := move.TargetSquare

	capturedPieceType := g.Board[moveTo].pieceType()
	movePiece := g.Board[moveFrom]
	movePieceType := movePiece.pieceType()

	isPromotion := move.flag == PromoteToQueen || move.flag == PromoteToKnight || move.flag == PromoteToRook || move.flag == PromoteToBishop
	isEnPassant := move.flag == EnPassantCapture

	currentGameState |= uint32(capturedPieceType) << 8
	if capturedPieceType != None && !isEnPassant {
		g.Board[moveTo] = Piece{None}
	}

	if movePieceType == King {
		if g.ColorToMove == White {
			newCastleState &= whiteCastleMask
		} else {
			newCastleState &= blackCastleMask
		}
	}

	if isPromotion {
		promoteType := 0
		switch move.flag {
		case PromoteToQueen:
			promoteType = Queen
		case PromoteToKnight:
			promoteType = Knight
		case PromoteToRook:
			promoteType = Rook
		case PromoteToBishop:
			promoteType = Bishop
		}
		g.Board[moveTo] = Piece{promoteType | g.ColorToMove}
	}
	if isEnPassant {
		epPawnSquare := 0
		if g.ColorToMove == White {
			epPawnSquare = moveTo - 8
		} else {
			epPawnSquare = moveTo + 8
		}
		currentGameState |= uint32((g.Board[epPawnSquare]).pieceType()) << 8 // add pawn as capture type
		g.Board[epPawnSquare] = Piece{None}                                  // clear en passant square
	}
	if move.flag == Castling {
		kingside := false
		if g.ColorToMove == White {
			kingside = moveTo == fromChessNotation("g1")
		} else {
			kingside = moveTo == fromChessNotation("g8")
		}
		castlingRookFromIndex := 0
		castlingRookToIndex := 0
		if kingside {
			castlingRookFromIndex = moveTo + 1
			castlingRookToIndex = moveTo - 1
		} else {
			castlingRookFromIndex = moveTo - 2
			castlingRookToIndex = moveTo + 1
		}
		g.Board[castlingRookFromIndex] = Piece{None}
		g.Board[castlingRookToIndex] = Piece{Rook | g.ColorToMove}
	}

	g.Board[moveTo] = movePiece
	g.Board[moveFrom] = Piece{None}

	if move.flag == PawnTwoForward {
		enPassantFile := uint32(moveFrom) % 8
		currentGameState |= (enPassantFile + 1) << 4
	}

	// If a piece moves to/from rook square, remove castling rights for that side
	if originalCastleRights != 0 {
		if moveTo == fromChessNotation("h1") || moveFrom == fromChessNotation("h1") {
			newCastleState &= whiteCastleKingsideMask
		}
		if moveTo == fromChessNotation("a1") || moveFrom == fromChessNotation("a1") {
			newCastleState &= whiteCastleQueensideMask
		}
		if moveTo == fromChessNotation("h8") || moveFrom == fromChessNotation("h8") {
			newCastleState &= blackCastleKingsideMask
		}
		if moveTo == fromChessNotation("a8") || moveFrom == fromChessNotation("a8") {
			newCastleState &= blackCastleQueensideMask
		}
	}

	currentGameState |= newCastleState
	currentGameState |= g.fiftyMoveCounter << 14
	g.currentGameState = currentGameState

	g.gameStateHistory = append(g.gameStateHistory, currentGameState)

	if g.ColorToMove == Black {
		g.plyCount++
		g.ColorToMove = White
	} else {
		g.ColorToMove = Black
	}
	g.fiftyMoveCounter++

	if movePieceType == Pawn || capturedPieceType != None {
		g.fiftyMoveCounter = 0
	}

	// piece := g.Board[move.StartSquare]
	// if piece.pieceType() == None {
	// 	return fmt.Errorf("no piece at %d", move.StartSquare)
	// }
	// if piece.color() != g.ColorToMove {
	// 	return fmt.Errorf("wrong color")
	// }
	// flag, err := g.movePiece(move, piece)
	// if err != nil {
	// 	return err
	// }

	// if g.ColorToMove == White {
	// 	g.ColorToMove = Black
	// } else {
	// 	g.ColorToMove = White
	// 	g.fullMoveCounter++
	// }

	// g.lastMove = Move{move.StartSquare, move.TargetSquare, flag}

	return nil
}

// func (g *Game) movePiece(move Move, p Piece) (int, error) {
// 	moves := g.LegalMoves(move.StartSquare)
// 	validMove := false
// 	flag := NoFlag
// 	for _, m := range moves {
// 		if m.TargetSquare == move.TargetSquare {
// 			flag = m.flag
// 			validMove = true
// 			break
// 		}
// 	}
// 	if !validMove {
// 		return flag, fmt.Errorf("invalid move")
// 	}

// 	if g.Board[move.TargetSquare].pieceType() == None && p.pieceType() != Pawn {
// 		currentFiftyMoveCounter := g.currentGameState >> 14 & 0x3f
// 		updatedFiftyMoveCounter := currentFiftyMoveCounter + 1
// 		g.currentGameState = g.currentGameState&^0x3f000 | updatedFiftyMoveCounter<<14
// 	}

// 	if flag == Castling {
// 		g.handleCastling(move)
// 	}

// 	if flag == EnPassantCapture {
// 		g.handleEnPassantCapture(move)
// 	}

// 	g.Board[move.TargetSquare] = p
// 	g.Board[move.StartSquare] = Piece{None}

// 	return flag, nil
// }

// func (g *Game) handleCastling(move Move) {
// 	if move.TargetSquare == 6 {
// 		g.Board[5] = g.Board[7]
// 		g.Board[7] = Piece{None}
// 	}
// 	if move.TargetSquare == 2 {
// 		g.Board[3] = g.Board[0]
// 		g.Board[0] = Piece{None}
// 	}
// 	if move.TargetSquare == 62 {
// 		g.Board[61] = g.Board[63]
// 		g.Board[63] = Piece{None}
// 	}
// 	if move.TargetSquare == 58 {
// 		g.Board[59] = g.Board[56]
// 		g.Board[56] = Piece{None}
// 	}
// }

// func (g *Game) handleEnPassantCapture(move Move) {
// 	if g.ColorToMove == White {
// 		g.Board[move.TargetSquare-8] = Piece{None}
// 	} else {
// 		g.Board[move.TargetSquare+8] = Piece{None}
// 	}
// }
