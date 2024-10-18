package game

import (
	"fmt"
	"strconv"
	"strings"
)

func (g *Game) loadPositionFromFen(fen string) error {
	pieces, color, castlingRights, enPassantSquare, fiftyMoveCounter, plyCount, err := parseFen(fen)
	if err != nil {
		return err
	}

	g.LoadPiecesFromFen(pieces)

	if color == "w" {
		g.ColorToMove = White
	} else if color == "b" {
		g.ColorToMove = Black
	} else {
		return fmt.Errorf("invalid color to move")
	}

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
		enPassantIndex := fromChessNotation(enPassantSquare)
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
		g.Board[rank*BoardSize+file] = createPiece(char)
		file++
	}
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
	if g.ColorToMove == White {
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
		epRank := 5
		if g.ColorToMove == Black {
			epRank = 2
		}
		fen += toChessNotation(epRank*BoardSize + int(enPassantFile-1))
	}

	fen += " "
	fen += strconv.Itoa(int(g.fiftyMoveCounter))

	fen += " "
	fen += strconv.Itoa(int(g.plyCount))

	return fen
}
