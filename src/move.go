package game

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

func (g *Game) generateMoves() []Move {
	moves := []Move{}

	for startSquare := 0; startSquare < BoardSize*BoardSize; startSquare++ {
		piece := g.Board[startSquare]
		if piece.pieceType() == None {
			continue
		}
		if piece.color() != g.ColorToMove {
			continue
		}

		if piece.pieceType() == Rook || piece.pieceType() == Bishop || piece.pieceType() == Queen {
			slidingMoves := g.generateSlidingMoves(startSquare)
			moves = append(moves, slidingMoves...)
		}
		if piece.pieceType() == Knight {
			knightMoves := g.generateKnightMoves(startSquare)
			moves = append(moves, knightMoves...)
		}
		if piece.pieceType() == King {
			kingMoves := g.generateKingMoves(startSquare)
			moves = append(moves, kingMoves...)
		}
		if piece.pieceType() == Pawn {
			pawnMoves := g.generatePawnMoves(startSquare)
			moves = append(moves, pawnMoves...)
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

			moves = append(moves, Move{startSquare, targetSquare})

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

		moves = append(moves, Move{startSquare, targetSquare})

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

		moves = append(moves, Move{startSquare, targetSquare})
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
		moves = append(moves, Move{startSquare, targetSquare})
		onFirstRank := startSquare/BoardSize == 1 && piece.color() == White || startSquare/BoardSize == 6 && piece.color() == Black
		if onFirstRank {
			targetSquare = startSquare + 16*direction
			if g.Board[targetSquare].pieceType() == None {
				moves = append(moves, Move{startSquare, targetSquare})
			}
		}
	}
	for _, offset := range []int{7, 9} {
		targetSquare = startSquare + offset*direction
		if g.Board[targetSquare].color() != piece.color() && g.Board[targetSquare].color() != None {
			moves = append(moves, Move{startSquare, targetSquare})
		}
	}

	// TODO: En passant
	return moves
}
