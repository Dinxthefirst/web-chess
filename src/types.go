package game

const BoardSize = 8

type Game struct {
	Board       [BoardSize][BoardSize]Piece
	ActiveColor Color
}

type Color int

const (
	White Color = iota
	Black
)

type Piece interface {
	Move(fromRow, fromCol, toRow, toCol int, g *Game) error
	Type() string
	Color() Color
	Symbol() string
}

var PieceUnicode = map[string]map[Color]string{
	"Pawn":   {White: "\u2659", Black: "\u265F"},
	"Rook":   {White: "\u2656", Black: "\u265C"},
	"Knight": {White: "\u2658", Black: "\u265E"},
	"Bishop": {White: "\u2657", Black: "\u265D"},
	"Queen":  {White: "\u2655", Black: "\u265B"},
	"King":   {White: "\u2654", Black: "\u265A"},
}

type GameState struct {
	Board       [BoardSize][BoardSize]PieceJSON `json:"board"`
	ActiveColor Color                           `json:"activeColor"`
}

type PieceJSON struct {
	Type   string `json:"type"`
	Color  Color  `json:"color"`
	Symbol string `json:"symbol"`
}

func (g *Game) ToGameState() GameState {
	var board [BoardSize][BoardSize]PieceJSON

	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			if piece := g.Board[x][y]; piece != nil {
				board[x][y] = PieceJSON{
					Type:   piece.Type(),
					Color:  piece.Color(),
					Symbol: piece.Symbol(),
				}
			}
		}
	}

	return GameState{
		Board:       board,
		ActiveColor: g.ActiveColor,
	}
}
