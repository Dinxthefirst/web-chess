package perft

import (
	"fmt"
	"time"
	game "web-chess/src"
)

func perft(g *game.Game, depth int) uint64 {
	if depth == 0 {
		return 1
	}

	moves := g.GenerateLegalMoves()
	var numPositions uint64 = 0

	for _, move := range moves {
		g.MakeMove(move)
		numPositions += perft(g, depth-1)
		g.UnmakeMove(move)
	}

	return numPositions
}

// RESULT:
// Depth: 0, Result: 1, Time: 0s
// Depth: 1, Result: 20, Time: 0s
// Depth: 2, Result: 400, Time: 1.0382ms
// Depth: 3, Result: 8906, Time: 21.2538ms
// Depth: 4, Result: 197530, Time: 442.9507ms
// Depth: 5, Result: 4877322, Time: 10.8970973s
// Depth: 6, Result: 119548601, Time: 5m9.07436s
// ACTUAL:
// https://www.chessprogramming.org/Perft_Results#Initial_Position
func RunPerft(depth int) {
	depths := []int{}
	for i := 0; i <= depth; i++ {
		depths = append(depths, i)
	}

	for _, d := range depths {
		start := time.Now()
		g := game.NewGame()
		numPositions := perft(g, d)
		fmt.Printf("Depth: %d, Result: %d, Time: %v\n", d, numPositions, time.Since(start))
	}
}
