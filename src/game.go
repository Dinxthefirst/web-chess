package game

const startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func NewGame() *Game {
	return NewGameFromFen(startFen)
}

func NewGameFromFen(fen string) *Game {
	precomputedMoveData()
	g := &Game{}
	g.loadPositionFromFen(fen)
	return g
}
