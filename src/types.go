package game

const BoardSize = 8

type Color int

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

type Game struct {
	Board       [BoardSize * BoardSize]Piece `json:"board"`
	ColorToMove Color                        `json:"ColorToMove"`
	currentFen  string
}

func (g *Game) GetFen() string {
	return g.currentFen
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
	flag         int
}

type Piece struct {
	Type int `json:"type"`
}

func (p *Piece) color() Color {
	return Color(p.Type & 24)
}

func (p *Piece) pieceType() int {
	return p.Type & 7
}
