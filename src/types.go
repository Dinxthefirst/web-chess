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
	ColorToMove Color                        `json:"activeColor"`
}

type Move struct {
	StartSquare  int `json:"startSquare"`
	TargetSquare int `json:"targetSquare"`
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
