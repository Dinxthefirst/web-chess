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
	var currentGameState uint32 = 0
	originalCastleRights := g.currentGameState & 15
	newCastleState := originalCastleRights

	moveFrom := move.StartSquare
	moveTo := move.TargetSquare

	capturedPieceType := g.Board[moveTo].pieceType()
	movePiece := g.Board[moveFrom]
	movePieceType := movePiece.pieceType()

	isPromotion := move.Flag == PromoteToQueen || move.Flag == PromoteToKnight || move.Flag == PromoteToRook || move.Flag == PromoteToBishop

	currentGameState |= uint32(capturedPieceType) << 8
	if capturedPieceType != None && move.Flag != EnPassantCapture {
		g.Board[moveTo] = Piece{None}
	}
	g.Board[moveTo] = movePiece

	if movePieceType == King {
		if g.ColorToMove == White {
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
		g.Board[moveTo] = Piece{promoteType | g.ColorToMove}
	} else if move.Flag == EnPassantCapture {
		epPawnSquare := 0
		if g.ColorToMove == White {
			epPawnSquare = moveTo - 8
		} else {
			epPawnSquare = moveTo + 8
		}
		currentGameState |= uint32((g.Board[epPawnSquare]).pieceType()) << 8 // add pawn as capture type
		g.Board[epPawnSquare] = Piece{None}                                  // clear en passant square
	} else if move.Flag == Castling {
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

	g.Board[moveFrom] = Piece{None}

	if move.Flag == PawnTwoForward {
		enPassantFile := uint32(moveFrom) % 8
		currentGameState |= (enPassantFile + 1) << 4
	}

	// If a piece moves to/from rook square, remove castling rights for that side
	if originalCastleRights != 0 {
		if moveTo == fromChessNotation("h1") || moveFrom == fromChessNotation("h1") {
			newCastleState &= whiteCastleKingsideMask
		} else if moveTo == fromChessNotation("a1") || moveFrom == fromChessNotation("a1") {
			newCastleState &= whiteCastleQueensideMask
		} else if moveTo == fromChessNotation("h8") || moveFrom == fromChessNotation("h8") {
			newCastleState &= blackCastleKingsideMask
		} else if moveTo == fromChessNotation("a8") || moveFrom == fromChessNotation("a8") {
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
}

func (g *Game) UnmakeMove(move Move) {
	opponentColor := g.ColorToMove
	g.ColorToMove = g.ColorToMove ^ 0b11000 // color is the color that made the move

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
	}

	if capturedPieceType != None && move.Flag != EnPassantCapture {
		g.Board[movedTo] = capturedPiece
	}

	g.Board[movedFrom] = Piece{movedPieceType | g.ColorToMove}
	g.Board[movedTo] = capturedPiece

	if move.Flag == EnPassantCapture {
		epPawnSquare := 0
		if g.ColorToMove == White {
			epPawnSquare = movedTo - 8
		} else {
			epPawnSquare = movedTo + 8
		}
		g.Board[movedTo] = Piece{None}
		g.Board[epPawnSquare] = capturedPiece
	} else if move.Flag == Castling {
		kingside := false
		if g.ColorToMove == White {
			kingside = movedTo == fromChessNotation("g1")
		} else {
			kingside = movedTo == fromChessNotation("g8")
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
		g.Board[castlingRookFromIndex] = Piece{Rook | g.ColorToMove}
	}

	g.gameStateHistory = g.gameStateHistory[:len(g.gameStateHistory)-1]
	currentGameState := g.gameStateHistory[len(g.gameStateHistory)-1]

	g.fiftyMoveCounter = (currentGameState >> 14) & 0b111111

	if g.ColorToMove == Black {
		g.plyCount--
	}

	g.currentGameState = currentGameState
}
