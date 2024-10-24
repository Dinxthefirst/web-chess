package game

import (
	"fmt"
	"strconv"
	"strings"
	"web-chess/backend/util"
)

func (g *Game) loadPositionFromFen(fen string) error {
	pieces, color, castlingRights, enPassantSquare, fiftyMoveCounter, plyCount, err := parseFen(fen)
	if err != nil {
		return err
	}

	g.LoadPiecesFromFen(pieces)

	g.ColorToMove = color == "w"

	var currentGameState uint32 = 0
	var newCastleState uint32 = 0
	if strings.Contains(castlingRights, "K") {
		newCastleState |= 0b1000
	}
	if strings.Contains(castlingRights, "Q") {
		newCastleState |= 0b0100
	}
	if strings.Contains(castlingRights, "k") {
		newCastleState |= 0b0010
	}
	if strings.Contains(castlingRights, "q") {
		newCastleState |= 0b0001
	}
	currentGameState |= newCastleState

	if enPassantSquare != "-" {
		enPassantIndex := util.FromChessNotation(enPassantSquare)
		enPassantFile := enPassantIndex % BoardSize
		currentGameState |= uint32(enPassantFile+1) << 4
	}

	g.currentGameState = currentGameState
	g.fiftyMoveCounter = uint32(fiftyMoveCounter)
	g.plyCount = plyCount

	g.gameStateHistory = []uint32{}
	g.gameStateHistory = append(g.gameStateHistory, currentGameState)
	return nil
}

func (g *Game) LoadPiecesFromFen(fen string) {
	rank := 7
	file := 0
	for _, char := range fen {
		if char == '/' {
			file = 0
			rank--
			continue
		}
		if char >= '1' && char <= '8' {
			file += int(char - '0')
			continue
		}
		index := rank*BoardSize + file
		piece := createPiece(char)
		g.Board[index] = piece
		g.updateBitboards(index, piece)
		file++
	}
	// print bitboards
	// fmt.Println("White Bitboard: ")
	// fmt.Println(bitboardString(g.whitePiecesBitBoard))
	// fmt.Println("Black Bitboard: ")
	// fmt.Println(bitboardString(g.blackPiecesBitBoard))
	// fmt.Println("Kings Bitboard: ")
	// fmt.Println(bitboardString(g.kingsBitBoard))
	// fmt.Println("Pawns Bitboard: ")
	// fmt.Println(bitboardString(g.pawnsBitBoard))
	// fmt.Println("Knights Bitboard: ")
	// fmt.Println(bitboardString(g.knightsBitBoard))
	// fmt.Println("Bishops Bitboard: ")
	// fmt.Println(bitboardString(g.bishopsBitBoard))
	// fmt.Println("Rooks Bitboard: ")
	// fmt.Println(bitboardString(g.rooksBitBoard))
	// fmt.Println("Queens Bitboard: ")
	// fmt.Println(bitboardString(g.queensBitBoard))
}

func (g *Game) updateBitboards(index int, piece Piece) {
	g.bitboards[piece.Type] |= 1 << index
}

func (g *Game) CurrentFen() string {
	fen := ""

	for rank := 7; rank >= 0; rank-- {
		emptySquares := 0

		for file := 0; file < 8; file++ {
			i := rank*BoardSize + file
			piece := g.Board[i]
			if piece.pieceType() == None {
				emptySquares++
				continue
			}
			if emptySquares > 0 {
				fen += fmt.Sprintf("%d", emptySquares)
				emptySquares = 0
			}
			fen += symbolForPiece(piece)
		}

		if emptySquares > 0 {
			fen += fmt.Sprintf("%d", emptySquares)
		}

		fen += "/"
	}
	fen = strings.TrimSuffix(fen, "/")

	fen += " "
	if g.ColorToMove {
		fen += "w"
	} else {
		fen += "b"
	}

	fen += " "
	currentCastlingRights := g.currentGameState & 0b1111
	castlingRights := ""
	if currentCastlingRights&0b1000 != 0 {
		castlingRights += "K"
	}
	if currentCastlingRights&0b0100 != 0 {
		castlingRights += "Q"
	}
	if currentCastlingRights&0b0010 != 0 {
		castlingRights += "k"
	}
	if currentCastlingRights&0b0001 != 0 {
		castlingRights += "q"
	}
	if castlingRights == "" {
		castlingRights = "-"
	}
	fen += castlingRights

	fen += " "
	enPassantFile := (g.currentGameState >> 4) & 0b1111
	if enPassantFile == 0 {
		fen += "-"
	} else {
		enPassantRank := 5
		if !g.ColorToMove {
			enPassantRank = 2
		}
		fen += util.ToChessNotation(enPassantRank*BoardSize + int(enPassantFile-1))
	}

	fen += " "
	fen += strconv.Itoa(int(g.fiftyMoveCounter))

	fen += " "
	fen += strconv.Itoa(int(g.plyCount))

	return fen
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
