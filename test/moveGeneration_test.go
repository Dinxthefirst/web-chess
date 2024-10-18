package test

import (
	game "web-chess/src"
)

func MoveGenerationTest(g *game.Game, depth int) int {
	if depth == 0 {
		return 1
	}

	moves := g.GenerateMoves()
	numPositions := 0

	for _, move := range moves {
		g.MakeMove(move)
		numPositions += MoveGenerationTest(g, depth-1)
		g.UnmakeMove(move)
	}

	return numPositions
}

// RESULT:
// Depth: 1, Result: 20, Time: 1.5699ms
// Depth: 2, Result: 400, Time: 25.4694ms
// Depth: 3, Result: 8906, Time: 551.2869ms
// Depth: 4, Result: 198008, Time: 13.0367955s
// Depth: 5, Result: 4909255, Time: 5m34.8636681s
// ACTUAL:
// 0 	1
// 1 	20
// 2 	400
// 3 	8,902
// 4 	197,281
// 5 	4,865,609
// func TestMoveGenerationFromStart(t *testing.T) {
// 	depth := []int{1, 2, 3, 4, 5}

// 	for _, d := range depth {
// 		start := time.Now()
// 		g := game.NewGame()
// 		numPositions := MoveGenerationTest(g, d)
// 		fmt.Printf("Depth: %d, Result: %d, Time: %v\n", d, numPositions, time.Since(start))
// 	}
// }
