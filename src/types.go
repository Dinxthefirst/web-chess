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
	Board       [BoardSize][BoardSize]Piece `json:"board"`
	ActiveColor Color                       `json:"activeColor"`
}

type Move struct {
	FromRow int `json:"fromRow"`
	FromCol int `json:"fromCol"`
	ToRow   int `json:"toRow"`
	ToCol   int `json:"toCol"`
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
