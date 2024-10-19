package game

const BoardSize = 8

const (
	None = iota
	King
	Pawn
	Knight
	Bishop
	Rook
	Queen
	White = 8
	Black = 16
)

const (
	whiteCastleKingsideMask  uint32 = 0b1111111111110111
	whiteCastleQueensideMask uint32 = 0b1111111111111011
	blackCastleKingsideMask  uint32 = 0b1111111111111101
	blackCastleQueensideMask uint32 = 0b1111111111111110

	whiteCastleMask uint32 = whiteCastleKingsideMask & whiteCastleQueensideMask
	blackCastleMask uint32 = blackCastleKingsideMask & blackCastleQueensideMask
)

type Game struct {
	Board [BoardSize * BoardSize]Piece `json:"board"`
	// 1: white, 0: black
	ColorToMove         bool `json:"ColorToMove"`
	kingsBitBoard       uint64
	pawnsBitBoard       uint64
	knightsBitBoard     uint64
	bishopsBitBoard     uint64
	rooksBitBoard       uint64
	queensBitBoard      uint64
	whitePiecesBitBoard uint64
	blackPiecesBitBoard uint64
	// Bits 0-3: white and black kingside/queen side castling rights
	//
	// Bit 0: black queenside,
	// Bit 1: black kingside,
	// Bit 2: white queenside,
	// Bit 3: white kingside
	//
	// Bits 4-7: file of en passant square (starting from 1, 0 means no en passant square)
	//
	// Bits 8-13: captured piece type
	//
	// Bits 14-19: fifty move counter
	//
	// Bits 20-31: full move counter
	currentGameState uint32
	gameStateHistory []uint32
	fiftyMoveCounter uint32
	plyCount         uint32
}

func (g *Game) BitBoards() [8]uint64 {
	return [8]uint64{g.kingsBitBoard, g.pawnsBitBoard, g.knightsBitBoard, g.bishopsBitBoard, g.rooksBitBoard, g.queensBitBoard, g.whitePiecesBitBoard, g.blackPiecesBitBoard}
}

const (
	NoFlag = iota
	EnPassantCapture
	Castling
	PromoteToQueen
	PromoteToKnight
	PromoteToRook
	PromoteToBishop
	PawnTwoForward
)

type Move struct {
	StartSquare  int `json:"startSquare"`
	TargetSquare int `json:"targetSquare"`
	Flag         int `json:"flag"`
}

type Piece struct {
	Type int `json:"type"`
}

func (p *Piece) color() int {
	return p.Type & 24
}

func (p *Piece) pieceType() int {
	return p.Type & 7
}
