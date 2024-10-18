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

// https://www.chessprogramming.org/Perft_Results#Initial_Position
func RunPerft(position, depth int) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	if position == 2 {
		fen = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
	} else if position == 3 {
		fen = "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1"
	} else if position == 4 {
		fen = "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1"
	} else if position == 5 {
		fen = "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8"
	} else if position == 6 {
		fen = "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10"
	}

	fmt.Printf("Running perft with depth %d on position %s\n", depth, fen)

	depths := []int{}
	for i := 0; i <= depth; i++ {
		depths = append(depths, i)
	}

	for _, d := range depths {
		start := time.Now()
		g := game.NewGameFromFen(fen)
		numPositions := perft(g, d)
		fmt.Printf("Depth: %d, Result: %d, Time: %v\n", d, numPositions, time.Since(start))
	}
}
