package game

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func fromChessNotation(notation string) int {
	if notation == "-" {
		return -1
	}
	file := int(notation[0] - 'a')
	rank := int(notation[1] - '1')
	return rank*BoardSize + file
}

func toChessNotation(square int) string {
	if square == -1 {
		return "-"
	}
	file := square % BoardSize
	rank := square / BoardSize
	return string(rune(file+'a')) + string(rune(rank+'1'))
}

func symbolForPiece(piece Piece) string {
	color := piece.color()
	pieceType := piece.pieceType()
	if color == White {
		switch pieceType {
		case Pawn:
			return "P"
		case Knight:
			return "N"
		case Bishop:
			return "B"
		case Rook:
			return "R"
		case Queen:
			return "Q"
		case King:
			return "K"
		}
	} else {
		switch pieceType {
		case Pawn:
			return "p"
		case Knight:
			return "n"
		case Bishop:
			return "b"
		case Rook:
			return "r"
		case Queen:
			return "q"
		case King:
			return "k"
		}
	}
	return ""
}
