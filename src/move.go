package game

var DirectionOffsets = [BoardSize]int{8, -8, 1, -1, 7, -7, 9, -9}

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
	}
	return moves
}

func (g *Game) generateSlidingMoves(startSquare int) []Move {
	moves := []Move{}
	piece := g.Board[startSquare]
	for directionIndex := 0; directionIndex < 8; directionIndex++ {
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
