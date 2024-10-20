package game

import "fmt"

func (g *Game) Move(move Move) error {
	if g.Board[move.StartSquare].pieceType() == None {
		return fmt.Errorf("no piece at %d", move.StartSquare)
	}

	moves := g.GenerateLegalMoves()
	validMove := false
	for _, m := range moves {
		if m.StartSquare == move.StartSquare && m.TargetSquare == move.TargetSquare {
			validMove = true
			if move.Flag == 0 {
				move.Flag = m.Flag
			}
			break
		}
	}
	if !validMove {
		return fmt.Errorf("no move from %d to %d", move.StartSquare, move.TargetSquare)
	}

	g.MakeMove(move)
	return nil
}

func (g *Game) MakeMove(move Move) {
	g.updateBitboardsWithMove(move)
	var currentGameState uint32 = 0
	originalCastleRights := g.currentGameState & 15
	newCastleState := originalCastleRights

	moveFrom := move.StartSquare
	moveTo := move.TargetSquare

	colorToMove := White
	if !g.ColorToMove {
		colorToMove = Black
	}

	capturedPieceType := g.Board[moveTo].pieceType()
	movePiece := g.Board[moveFrom]
	originalPieceType := movePiece.pieceType()
	movePieceType := movePiece.pieceType()

	isPromotion := move.Flag == PromoteToQueen || move.Flag == PromoteToKnight || move.Flag == PromoteToRook || move.Flag == PromoteToBishop

	currentGameState |= uint32(capturedPieceType) << 8
	// if capturedPieceType != None && move.Flag != EnPassantCapture {
	// 	g.Board[moveTo] = Piece{None}
	// }

	if movePieceType == King {
		if g.ColorToMove {
			newCastleState &= whiteCastleMask
		} else {
			newCastleState &= blackCastleMask
		}
	}

	if isPromotion {
		promoteType := 0
		switch move.Flag {
		case PromoteToQueen:
			promoteType = Queen
		case PromoteToKnight:
			promoteType = Knight
		case PromoteToRook:
			promoteType = Rook
		case PromoteToBishop:
			promoteType = Bishop
		}
		movePieceType = promoteType
		// g.pawnsBitBoard &= ^(1 << moveFrom)
		movePiece = Piece{promoteType | colorToMove}
	} else if move.Flag == EnPassantCapture {
		epPawnSquare := 0
		if g.ColorToMove {
			epPawnSquare = moveTo - 8
		} else {
			epPawnSquare = moveTo + 8
		}
		currentGameState |= uint32((g.Board[epPawnSquare]).pieceType()) << 8 // add pawn as capture type
		g.Board[epPawnSquare] = Piece{None}                                  // clear en passant square
		// fmt.Printf("Pawns Bitboard before:\n%s", bitboardString(g.pawnsBitBoard))
		// fmt.Printf("En Passant mask:\n%s", bitboardString(^(1 << epPawnSquare)))
		// g.pawnsBitBoard &= ^(1 << epPawnSquare) // Removes captured pawn
		// g.pawnsBitBoard &= ^(1 << moveFrom)     // Removes moving pawn from original square
		// g.pawnsBitBoard |= 1 << moveTo          // Adds moving pawn to new square
		// if g.ColorToMove {                      // white moved and captured black pawn
		// 	g.blackPiecesBitBoard &= ^(1 << epPawnSquare)
		// 	g.whitePiecesBitBoard &= ^(1 << moveFrom)
		// 	g.whitePiecesBitBoard |= 1 << moveTo
		// } else {
		// 	g.whitePiecesBitBoard &= ^(1 << epPawnSquare)
		// 	g.blackPiecesBitBoard &= ^(1 << moveFrom)
		// 	g.blackPiecesBitBoard |= 1 << moveTo
		// }
		// fmt.Printf("Pawns Bitboard after:\n%s", bitboardString(g.pawnsBitBoard))
	} else if move.Flag == Castling {
		kingside := false
		if g.ColorToMove {
			kingside = moveTo == 6 // g1
		} else {
			kingside = moveTo == 62 // g8
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
		g.Board[castlingRookToIndex] = Piece{Rook | colorToMove}
		// g.rooksBitBoard &= ^(1 << castlingRookFromIndex)
		// g.rooksBitBoard |= 1 << castlingRookToIndex
		// if g.ColorToMove {
		// 	g.whitePiecesBitBoard &= ^(1 << castlingRookFromIndex)
		// 	g.whitePiecesBitBoard |= 1 << castlingRookToIndex
		// } else {
		// 	g.blackPiecesBitBoard &= ^(1 << castlingRookFromIndex)
		// 	g.blackPiecesBitBoard |= 1 << castlingRookToIndex
		// }
	}

	g.Board[moveTo] = movePiece
	g.Board[moveFrom] = Piece{None}

	// switch movePieceType {
	// case King:
	// 	g.kingsBitBoard &= ^(1 << moveFrom)
	// 	g.kingsBitBoard |= 1 << moveTo
	// case Pawn:
	// 	if move.Flag != EnPassantCapture {
	// 		g.pawnsBitBoard &= ^(1 << moveFrom)
	// 		g.pawnsBitBoard |= 1 << moveTo
	// 	}
	// case Knight:
	// 	g.knightsBitBoard &= ^(1 << moveFrom)
	// 	g.knightsBitBoard |= 1 << moveTo
	// case Bishop:
	// 	g.bishopsBitBoard &= ^(1 << moveFrom)
	// 	g.bishopsBitBoard |= 1 << moveTo
	// case Rook:
	// 	g.rooksBitBoard &= ^(1 << moveFrom)
	// 	g.rooksBitBoard |= 1 << moveTo
	// case Queen:
	// 	g.queensBitBoard &= ^(1 << moveFrom)
	// 	g.queensBitBoard |= 1 << moveTo
	// }
	// if move.Flag != EnPassantCapture {
	// 	if colorToMove == White {
	// 		g.whitePiecesBitBoard &= ^(1 << moveFrom)
	// 		g.whitePiecesBitBoard |= 1 << moveTo
	// 	} else {
	// 		g.blackPiecesBitBoard &= ^(1 << moveFrom)
	// 		g.blackPiecesBitBoard |= 1 << moveTo
	// 	}
	// }

	if move.Flag == PawnTwoForward {
		enPassantFile := uint32(moveFrom) % 8
		currentGameState |= (enPassantFile + 1) << 4
	}

	// If a piece moves to/from rook square, remove castling rights for that side
	if originalCastleRights != 0 {
		if moveTo == 7 || moveFrom == 7 { // h1
			newCastleState &= whiteCastleKingsideMask
		} else if moveTo == 0 || moveFrom == 0 { // a1
			newCastleState &= whiteCastleQueensideMask
		} else if moveTo == 63 || moveFrom == 63 { // h8
			newCastleState &= blackCastleKingsideMask
		} else if moveTo == 56 || moveFrom == 56 { // a8
			newCastleState &= blackCastleQueensideMask
		}
	}

	currentGameState |= newCastleState
	currentGameState |= g.fiftyMoveCounter << 14
	g.currentGameState = currentGameState

	g.gameStateHistory = append(g.gameStateHistory, currentGameState)

	if !g.ColorToMove {
		g.plyCount++
	}
	g.ColorToMove = !g.ColorToMove

	g.fiftyMoveCounter++

	if originalPieceType == Pawn || capturedPieceType != None {
		g.fiftyMoveCounter = 0
	}
}

func (g *Game) UnmakeMove(move Move) {
	g.ColorToMove = !g.ColorToMove // color is the color that made the move

	colorToMove := White
	opponentColor := Black
	if !g.ColorToMove {
		colorToMove = Black
		opponentColor = White
	}

	capturedPieceType := (g.currentGameState >> 8) & 0b111111
	capturedPiece := Piece{None}
	if capturedPieceType != None {
		capturedPiece = Piece{int(capturedPieceType) | opponentColor}
	}

	movedFrom := move.StartSquare
	movedTo := move.TargetSquare

	isPromotion := move.Flag == PromoteToQueen || move.Flag == PromoteToKnight || move.Flag == PromoteToRook || move.Flag == PromoteToBishop

	toSquarePieceType := g.Board[movedTo]
	movedPieceType := toSquarePieceType.pieceType()
	if isPromotion {
		movedPieceType = Pawn
		// switch move.Flag {
		// case PromoteToQueen:
		// 	g.queensBitBoard &= ^(1 << movedTo)
		// case PromoteToKnight:
		// 	g.knightsBitBoard &= ^(1 << movedTo)
		// case PromoteToRook:
		// 	g.rooksBitBoard &= ^(1 << movedTo)
		// case PromoteToBishop:
		// 	g.bishopsBitBoard &= ^(1 << movedTo)
		// }
	}

	// if capturedPieceType != None && move.Flag != EnPassantCapture {
	// 	g.Board[movedTo] = capturedPiece
	// }

	g.Board[movedFrom] = Piece{movedPieceType | colorToMove}
	g.Board[movedTo] = capturedPiece

	if move.Flag == EnPassantCapture {
		epPawnSquare := 0
		if g.ColorToMove {
			epPawnSquare = movedTo - 8
		} else {
			epPawnSquare = movedTo + 8
		}
		g.Board[movedTo] = Piece{None}
		g.Board[epPawnSquare] = capturedPiece
		// fmt.Printf("\nUnmaking en passant\n")
		// fmt.Printf("Pawns Bitboard:\n%s", bitboardString(g.pawnsBitBoard))
		// fmt.Printf("En Passant mask:\n%s", bitboardString(^(1 << epPawnSquare)))
		// g.pawnsBitBoard |= 1 << epPawnSquare // Adds captured pawn
		// g.pawnsBitBoard |= 1 << movedFrom    // Adds moving pawn back to original square
		// g.pawnsBitBoard &= ^(1 << movedTo)   // Removes moving pawn from new square
		// if g.ColorToMove {                   // white moved and captured black pawn
		// 	g.blackPiecesBitBoard |= 1 << epPawnSquare
		// 	g.whitePiecesBitBoard |= 1 << movedFrom
		// 	g.whitePiecesBitBoard &= ^(1 << movedTo)
		// } else {
		// 	g.whitePiecesBitBoard |= 1 << epPawnSquare
		// 	g.blackPiecesBitBoard |= 1 << movedFrom
		// 	g.blackPiecesBitBoard &= ^(1 << movedTo)
		// }
		// fmt.Printf("Pawns Bitboard after:\n%s", bitboardString(g.pawnsBitBoard))
	} else if move.Flag == Castling {
		kingside := false
		if g.ColorToMove {
			kingside = movedTo == 6 // g1
		} else {
			kingside = movedTo == 62 //
		}
		castlingRookFromIndex := 0
		castlingRookToIndex := 0
		if kingside {
			castlingRookFromIndex = movedTo + 1
			castlingRookToIndex = movedTo - 1
		} else {
			castlingRookFromIndex = movedTo - 2
			castlingRookToIndex = movedTo + 1
		}
		g.Board[castlingRookToIndex] = Piece{None}
		g.Board[castlingRookFromIndex] = Piece{Rook | colorToMove}
		// g.rooksBitBoard |= 1 << castlingRookFromIndex
		// g.rooksBitBoard &= ^(1 << castlingRookToIndex)
		// if g.ColorToMove {
		// 	g.whitePiecesBitBoard |= 1 << castlingRookFromIndex
		// 	g.whitePiecesBitBoard &= ^(1 << castlingRookToIndex)
		// } else {
		// 	g.blackPiecesBitBoard |= 1 << castlingRookFromIndex
		// 	g.blackPiecesBitBoard &= ^(1 << castlingRookToIndex)
		// }
	}

	// switch movedPieceType {
	// case King:
	// 	g.kingsBitBoard |= 1 << movedFrom
	// 	g.kingsBitBoard &= ^(1 << movedTo)
	// case Pawn:
	// 	if move.Flag != EnPassantCapture {
	// 		g.pawnsBitBoard |= 1 << movedFrom
	// 		g.pawnsBitBoard &= ^(1 << movedTo)
	// 	}
	// case Knight:
	// 	g.knightsBitBoard |= 1 << movedFrom
	// 	g.knightsBitBoard &= ^(1 << movedTo)
	// case Bishop:
	// 	g.bishopsBitBoard |= 1 << movedFrom
	// 	g.bishopsBitBoard &= ^(1 << movedTo)
	// case Rook:
	// 	g.rooksBitBoard |= 1 << movedFrom
	// 	g.rooksBitBoard &= ^(1 << movedTo)
	// case Queen:
	// 	g.queensBitBoard |= 1 << movedFrom
	// 	g.queensBitBoard &= ^(1 << movedTo)
	// }
	// if move.Flag != EnPassantCapture {
	// 	if colorToMove == White {
	// 		g.whitePiecesBitBoard |= 1 << movedFrom
	// 		g.whitePiecesBitBoard &= ^(1 << movedTo)
	// 	} else {
	// 		g.blackPiecesBitBoard |= 1 << movedFrom
	// 		g.blackPiecesBitBoard &= ^(1 << movedTo)
	// 	}
	// }
	// switch capturedPieceType {
	// case Pawn:
	// 	if move.Flag != EnPassantCapture {
	// 		g.pawnsBitBoard |= 1 << movedTo
	// 	}
	// case Knight:
	// 	g.knightsBitBoard |= 1 << movedTo
	// case Bishop:
	// 	g.bishopsBitBoard |= 1 << movedTo
	// case Rook:
	// 	g.rooksBitBoard |= 1 << movedTo
	// case Queen:
	// 	g.queensBitBoard |= 1 << movedTo
	// }
	// if capturedPieceType != None {
	// 	if colorToMove == White {
	// 		g.blackPiecesBitBoard |= 1 << movedTo
	// 	} else {
	// 		g.whitePiecesBitBoard |= 1 << movedTo
	// 	}
	// }

	g.gameStateHistory = g.gameStateHistory[:len(g.gameStateHistory)-1]
	currentGameState := g.gameStateHistory[len(g.gameStateHistory)-1]

	g.fiftyMoveCounter = (currentGameState >> 14) & 0b111111

	if !g.ColorToMove {
		g.plyCount--
	}

	g.currentGameState = currentGameState
}

func (g *Game) updateBitboardsWithMove(move Move) {
	pieceToMove := g.Board[move.StartSquare].Type
	pieceToCapture := g.Board[move.TargetSquare].Type

	g.bitboards[pieceToMove] ^= 1<<move.StartSquare | 1<<move.TargetSquare
	if pieceToCapture != None {
		g.bitboards[pieceToCapture] ^= 1 << move.TargetSquare
	}
	fmt.Println(g.BitBoards())
}

// func (g *Game) undoBitboardsWithMove(move Move) {
// 	pieceMoved := g.Board[move.TargetSquare].Type

// }
