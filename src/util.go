package game

import (
	"strconv"
	"strings"
)

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

func parseFen(fen string) (pieces, color, castlingRights, enPassantSquare string, fiftyMoveCounter, plyCount uint32, err error) {
	splitFen := strings.Split(fen, " ")
	pieces = splitFen[0]
	color = splitFen[1]
	castlingRights = splitFen[2]
	enPassantSquare = splitFen[3]
	fiftyMoveCounterInt, err := strconv.Atoi(splitFen[4])
	if err != nil {
		return "", "", "", "", 0, 0, err
	}
	fiftyMoveCounter = uint32(fiftyMoveCounterInt)

	plyCountInt, err := strconv.Atoi(splitFen[5])
	if err != nil {
		return "", "", "", "", 0, 0, err
	}
	plyCount = uint32(plyCountInt)
	return pieces, color, castlingRights, enPassantSquare, fiftyMoveCounter, plyCount, nil
}

func createPiece(char rune) Piece {
	color := White
	if char >= 'a' && char <= 'z' {
		color = Black
	}
	pieceType := None
	switch char {
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
	return Piece{pieceType | color}
}

// func bitboardString(bitboard uint64) string {
// 	str := ""
// 	for rank := 7; rank >= 0; rank-- {
// 		for file := 0; file < 8; file++ {
// 			str += fmt.Sprintf("%b", bitboard&(1<<(rank*BoardSize+file))>>(rank*BoardSize+file))
// 		}
// 		str += "\n"
// 	}
// 	return str
// }
