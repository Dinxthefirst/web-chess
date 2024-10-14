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

var PieceImages = map[string]map[Color]string{
	"Pawn":   {White: "/static/images/white_pawn.svg", Black: "/static/images/black_pawn.svg"},
	"Rook":   {White: "/static/images/white_rook.svg", Black: "/static/images/black_rook.svg"},
	"Knight": {White: "/static/images/white_knight.svg", Black: "/static/images/black_knight.svg"},
	"Bishop": {White: "/static/images/white_bishop.svg", Black: "/static/images/black_bishop.svg"},
	"Queen":  {White: "/static/images/white_queen.svg", Black: "/static/images/black_queen.svg"},
	"King":   {White: "/static/images/white_king.svg", Black: "/static/images/black_king.svg"},
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
