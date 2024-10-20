package game

import (
	"fmt"
)

var DirectionOffsets = [BoardSize]int{8, -8, 1, -1, 7, -7, 9, -9}
var KnightOffsets = [8]int{15, 17, 10, 6, -15, -17, -10, -6}

var NumSquaresToEdge [BoardSize * BoardSize][8]int

// var opponentAttackMap uint64
// var opponentAttackMapSliding uint64
var whiteKingsideAttacked bool
var whiteQueensideAttacked bool
var blackKingsideAttacked bool
var blackQueensideAttacked bool

func precomputedMoveData() {
	for file := 0; file < BoardSize; file++ {
		for rank := 0; rank < BoardSize; rank++ {
			numNorth := (BoardSize - 1) - rank
			numSouth := rank
			numEast := (BoardSize - 1) - file
			numWest := file

			squareIndex := rank*BoardSize + file
			NumSquaresToEdge[squareIndex] = [8]int{
				numNorth,
				numSouth,
				numEast,
				numWest,
				min(numNorth, numWest),
				min(numSouth, numEast),
				min(numNorth, numEast),
				min(numSouth, numWest),
			}
		}
	}
}

func (g *Game) LegalMovesAtIndex(index int) []Move {
	moves := g.GenerateLegalMoves()

	filteredMoves := []Move{}
	for _, move := range moves {
		if move.StartSquare == index {
			filteredMoves = append(filteredMoves, move)
		}
	}
	return filteredMoves
}

func (g *Game) GenerateLegalMoves() []Move {
	moves := g.GeneratePseudoLegalMoves()

	filteredMoves := []Move{}
	for _, move := range moves {
		if g.isMoveLegal(move) {
			filteredMoves = append(filteredMoves, move)
		}
	}

	return filteredMoves
}

func (g *Game) isMoveLegal(move Move) bool {
	color := g.ColorToMove

	g.MakeMove(move)

	isKingInCheck := g.isKingInCheck(color)

	g.UnmakeMove(move)

	return !isKingInCheck
}

func (g *Game) isKingInCheck(color bool) bool {
	kingPosition := g.findKing(color)
	return g.isSquareAttacked(kingPosition, color)
}

func (g *Game) findKing(color bool) int {
	for i, piece := range g.Board {
		if piece.pieceType() != King {
			continue
		}
		pieceColor := piece.color()
		if pieceColor == White && color || pieceColor == Black && !color {
			return i
		}
	}
	return -1 // Should never happen
}

func (g *Game) isSquareAttacked(square int, color bool) bool {
	opponentColor := !color

	for _, move := range g.generateMovesForColor(opponentColor, true) {
		if move.TargetSquare == square {
			return true
		}
	}
	return false
}

func (g *Game) GeneratePseudoLegalMoves() []Move {
	return g.generateMovesForColor(g.ColorToMove, false)
}

func (g *Game) generateMovesForColor(color bool, inSearch bool) []Move {
	moves := []Move{}

	for startSquare := 0; startSquare < BoardSize*BoardSize; startSquare++ {
		piece := g.Board[startSquare]
		if piece.pieceType() == None {
			continue
		}
		if piece.color() == White && !color || piece.color() == Black && color {
			continue
		}

		switch piece.pieceType() {
		case King:
			moves = append(moves, g.generateKingMoves(startSquare, inSearch)...)
		case Pawn:
			moves = append(moves, g.generatePawnMoves(startSquare)...)
		case Knight:
			moves = append(moves, g.generateKnightMoves(startSquare)...)
		case Bishop, Rook, Queen:
			moves = append(moves, g.generateSlidingMoves(startSquare)...)
		}
	}
	return moves
}

func (g *Game) generateSlidingMoves(startSquare int) []Move {
	piece := g.Board[startSquare]

	startDirIndex := 0
	if piece.pieceType() == Bishop {
		startDirIndex = 4
	}
	endDirIndex := 8
	if piece.pieceType() == Rook {
		endDirIndex = 4
	}

	moves := []Move{}
	for directionIndex := startDirIndex; directionIndex < endDirIndex; directionIndex++ {
		for numSquares := 0; numSquares < NumSquaresToEdge[startSquare][directionIndex]; numSquares++ {
			targetSquare := startSquare + DirectionOffsets[directionIndex]*(numSquares+1)

			pieceOnTargetSquare := g.Board[targetSquare]

			if pieceOnTargetSquare.color() == piece.color() {
				break
			}

			moves = append(moves, Move{startSquare, targetSquare, NoFlag})

			if pieceOnTargetSquare.color() != None {
				break
			}
		}
	}
	return moves
}

func (g *Game) generateKnightMoves(startSquare int) []Move {
	piece := g.Board[startSquare]

	moves := []Move{}
	rank := startSquare / BoardSize
	file := startSquare % BoardSize
	for _, offset := range KnightOffsets {
		targetSquare := startSquare + offset

		if targetSquare < 0 || targetSquare >= BoardSize*BoardSize {
			continue
		}

		targetRank := targetSquare / BoardSize
		targetFile := targetSquare % BoardSize

		if abs(rank-targetRank) > 2 || abs(file-targetFile) > 2 {
			continue
		}

		pieceOnTargetSquare := g.Board[targetSquare]

		if pieceOnTargetSquare.color() == piece.color() {
			continue
		}

		moves = append(moves, Move{startSquare, targetSquare, NoFlag})

	}
	return moves
}

func (g *Game) generateKingMoves(startSquare int, inSearch bool) []Move {
	piece := g.Board[startSquare]

	moves := []Move{}
	for i, offset := range DirectionOffsets {
		if NumSquaresToEdge[startSquare][i] == 0 {
			continue
		}

		targetSquare := startSquare + offset

		pieceOnTargetSquare := g.Board[targetSquare]

		if pieceOnTargetSquare.color() == piece.color() {
			continue
		}

		moves = append(moves, Move{startSquare, targetSquare, NoFlag})
	}

	moves = append(moves, g.generateCastlingMoves(startSquare, inSearch)...)

	return moves
}

func (g *Game) generateCastlingMoves(startSquare int, inSearch bool) []Move {
	moves := []Move{}

	if inSearch {
		return moves
	}
	g.generateCastleAttackedSquares()

	if g.currentGameState&0b1111 == 0 {
		return moves
	}

	castlingRights := g.currentGameState & 0b1111
	if g.ColorToMove {
		if ((castlingRights >> 3) & 1) == 1 {
			bishopMoved := g.Board[fromChessNotation("f1")].pieceType() == None
			knightMoved := g.Board[fromChessNotation("g1")].pieceType() == None
			if bishopMoved && knightMoved {
				if !whiteKingsideAttacked {
					moves = append(moves, Move{startSquare, 6, Castling})
				}
			}
		}
		if ((castlingRights >> 2) & 1) == 1 {
			knightMoved := g.Board[fromChessNotation("b1")].pieceType() == None
			bishopMoved := g.Board[fromChessNotation("c1")].pieceType() == None
			queenMoved := g.Board[fromChessNotation("d1")].pieceType() == None
			if knightMoved && bishopMoved && queenMoved {
				if !whiteQueensideAttacked {
					moves = append(moves, Move{startSquare, 2, Castling})
				}
			}
		}
	} else {
		if ((castlingRights >> 1) & 1) == 1 {
			bishopMoved := g.Board[fromChessNotation("f8")].pieceType() == None
			knightMoved := g.Board[fromChessNotation("g8")].pieceType() == None
			if bishopMoved && knightMoved {
				if !blackKingsideAttacked {
					moves = append(moves, Move{startSquare, 62, Castling})
				}
			}
		}
		if (castlingRights & 1) == 1 {
			knightMoved := g.Board[fromChessNotation("b8")].pieceType() == None
			bishopMoved := g.Board[fromChessNotation("c8")].pieceType() == None
			queenMoved := g.Board[fromChessNotation("d8")].pieceType() == None
			if knightMoved && bishopMoved && queenMoved {
				if !blackQueensideAttacked {
					moves = append(moves, Move{startSquare, 58, Castling})
				}
			}
		}
	}

	return moves
}

func (g *Game) generateCastleAttackedSquares() {
	opponentColor := !g.ColorToMove

	whiteKingsideAttacked = false
	whiteQueensideAttacked = false
	blackKingsideAttacked = false
	blackQueensideAttacked = false

	for _, move := range g.generateMovesForColor(opponentColor, true) {
		if whiteKingsideAttacked && whiteQueensideAttacked && blackKingsideAttacked && blackQueensideAttacked {
			return
		}

		if move.TargetSquare == 5 {
			whiteKingsideAttacked = true
		} else if move.TargetSquare == 3 {
			whiteQueensideAttacked = true
		} else if move.TargetSquare == 61 {
			blackKingsideAttacked = true
		} else if move.TargetSquare == 59 {
			blackQueensideAttacked = true
		}
	}
}

func (g *Game) generatePawnMoves(startSquare int) []Move {
	piece := g.Board[startSquare]

	moves := []Move{}
	direction := 1
	if piece.color() == Black {
		direction = -1
	}

	targetSquare := startSquare + 8*direction
	if targetSquare < 0 || targetSquare >= BoardSize*BoardSize {
		fmt.Println("Start square: ", startSquare)
		fmt.Println(g.CurrentFen())
	}

	promotionMoveAllowed := startSquare/BoardSize == 6 && piece.color() == White || startSquare/BoardSize == 1 && piece.color() == Black
	if g.Board[targetSquare].pieceType() == None {
		if promotionMoveAllowed {
			moves = append(moves, Move{startSquare, targetSquare, PromoteToQueen})
			moves = append(moves, Move{startSquare, targetSquare, PromoteToKnight})
			moves = append(moves, Move{startSquare, targetSquare, PromoteToRook})
			moves = append(moves, Move{startSquare, targetSquare, PromoteToBishop})
		} else {
			moves = append(moves, Move{startSquare, targetSquare, NoFlag})
		}
		doubleMoveAllowed := startSquare/BoardSize == 1 && piece.color() == White || startSquare/BoardSize == 6 && piece.color() == Black
		if doubleMoveAllowed {
			targetSquare = startSquare + 16*direction
			if g.Board[targetSquare].pieceType() == None {
				moves = append(moves, Move{startSquare, targetSquare, PawnTwoForward})
			}
		}
	}

	for _, offset := range []int{7, 9} {
		targetSquare = startSquare + offset*direction
		if targetSquare < 0 || targetSquare >= BoardSize*BoardSize {
			continue
		}

		startFile := startSquare % BoardSize
		targetFile := targetSquare % BoardSize
		if abs(startFile-targetFile) != 1 {
			continue
		}

		if g.Board[targetSquare].color() != piece.color() && g.Board[targetSquare].color() != None {
			if promotionMoveAllowed {
				moves = append(moves, Move{startSquare, targetSquare, PromoteToQueen})
				moves = append(moves, Move{startSquare, targetSquare, PromoteToKnight})
				moves = append(moves, Move{startSquare, targetSquare, PromoteToRook})
				moves = append(moves, Move{startSquare, targetSquare, PromoteToBishop})
			} else {
				moves = append(moves, Move{startSquare, targetSquare, NoFlag})
			}
		}

		enPassantFile := g.currentGameState >> 4 & 0b111
		if enPassantFile == 0 {
			continue
		}
		enPassantRank := 5
		if piece.color() == Black {
			enPassantRank = 2
		}
		enPassantSquare := enPassantRank*BoardSize + int(enPassantFile-1)
		if targetSquare == enPassantSquare {
			moves = append(moves, Move{startSquare, targetSquare, EnPassantCapture})
		}
	}

	return moves
}
