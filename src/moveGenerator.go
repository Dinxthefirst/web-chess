package game

import (
	"strings"
)

var DirectionOffsets = [BoardSize]int{8, -8, 1, -1, 7, -7, 9, -9}
var KnightOffsets = [8]int{15, 17, 10, 6, -15, -17, -10, -6}

var NumSquaresToEdge [BoardSize * BoardSize][]int

func precomputedMoveData() {
	for file := 0; file < BoardSize; file++ {
		for rank := 0; rank < BoardSize; rank++ {
			numNorth := (BoardSize - 1) - rank
			numSouth := rank
			numEast := (BoardSize - 1) - file
			numWest := file

			squareIndex := rank*BoardSize + file
			NumSquaresToEdge[squareIndex] = []int{
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

func (g *Game) LegalMoves(index int) []Move {
	moves := g.generateMovesForColor(g.ColorToMove)

	filteredMoves := []Move{}
	for _, move := range moves {
		if move.StartSquare == index {
			filteredMoves = append(filteredMoves, move)
		}
	}
	return filteredMoves
}

func (g *Game) GenerateMoves() []Move {
	return g.generateMovesForColor(g.ColorToMove)
}

func (g *Game) generateMovesForColor(color int) []Move {
	moves := []Move{}

	for startSquare := 0; startSquare < BoardSize*BoardSize; startSquare++ {
		piece := g.Board[startSquare]
		if piece.pieceType() == None {
			continue
		}
		if piece.color() != color {
			continue
		}

		switch piece.pieceType() {
		case King:
			moves = append(moves, g.generateKingMoves(startSquare)...)
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

		targetRank := targetSquare / BoardSize
		targetFile := targetSquare % BoardSize

		if targetRank < 0 || targetRank >= BoardSize || targetFile < 0 || targetFile >= BoardSize {
			continue
		}

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

func (g *Game) generateKingMoves(startSquare int) []Move {
	piece := g.Board[startSquare]

	moves := []Move{}
	for _, offset := range DirectionOffsets {
		targetSquare := startSquare + offset

		if targetSquare < 0 || targetSquare >= BoardSize*BoardSize {
			continue
		}

		pieceOnTargetSquare := g.Board[targetSquare]

		if pieceOnTargetSquare.color() == piece.color() {
			continue
		}

		moves = append(moves, Move{startSquare, targetSquare, NoFlag})
	}

	moves = append(moves, g.generateCastlingMoves(startSquare)...)

	return moves
}

func (g *Game) generateCastlingMoves(startSquare int) []Move {
	moves := []Move{}

	if g.currentGameState&0b1111 == 0 {
		return moves
	}

	castlingRights := g.currentGameState & 0b1111
	if g.ColorToMove == White {
		if ((castlingRights >> 3) & 1) == 1 {
			moves = append(moves, Move{startSquare, 6, Castling})
		}
		if ((castlingRights >> 2) & 1) == 1 {
			moves = append(moves, Move{startSquare, 2, Castling})
		}
	} else {
		if ((castlingRights >> 1) & 1) == 1 {
			moves = append(moves, Move{startSquare, 62, Castling})
		}
		if (castlingRights & 1) == 1 {
			moves = append(moves, Move{startSquare, 58, Castling})
		}
	}

	return moves
}

func (g *Game) generatePawnMoves(startSquare int) []Move {
	piece := g.Board[startSquare]

	moves := []Move{}
	direction := 1
	if piece.color() == Black {
		direction = -1
	}

	targetSquare := startSquare + 8*direction
	if g.Board[targetSquare].pieceType() == None {
		moves = append(moves, Move{startSquare, targetSquare, NoFlag})
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
		if g.Board[targetSquare].color() != piece.color() && g.Board[targetSquare].color() != None {
			moves = append(moves, Move{startSquare, targetSquare, NoFlag})
		}
		splitFen := strings.Split(g.CurrentFen(), " ")
		enPassantSquare := fromChessNotation(splitFen[3])
		if enPassantSquare != -1 && targetSquare == enPassantSquare {
			moves = append(moves, Move{startSquare, targetSquare, EnPassantCapture})
		}
	}

	return moves
}
