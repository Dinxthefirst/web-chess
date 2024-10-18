package perft

import (
	"fmt"
	"sort"
	"time"
	game "web-chess/src"
)

var actualResults = map[int]map[int]uint64{
	1: {
		1: 20,
		2: 400,
		3: 8902,
		4: 197281,
		5: 4865609,
		6: 119060324,
		7: 3195901860,
		8: 84998978956,
	},
	2: {
		1: 48,
		2: 2039,
		3: 97862,
		4: 4085603,
		5: 193690690,
	},
	3: {
		1: 14,
		2: 191,
		3: 2812,
		4: 43238,
		5: 674624,
	},
	4: {
		1: 6,
		2: 264,
		3: 9467,
		4: 422333,
		5: 15833292,
	},
	5: {
		1: 44,
		2: 1486,
		3: 62379,
		4: 2103487,
		5: 89941194,
	},
	6: {
		1: 46,
		2: 2079,
		3: 89890,
		4: 3894594,
		5: 164075551,
	},
}

func perft(g *game.Game, depth int) uint64 {
	moves := g.GenerateLegalMoves()

	if depth == 1 {
		return uint64(len(moves))
	}
	var numPositions uint64 = 0

	for _, move := range moves {
		g.MakeMove(move)
		numPositions += perft(g, depth-1)
		g.UnmakeMove(move)
	}

	return numPositions
}

func RunPerftTest() {
	RunPerft(1, 4)
	RunPerft(2, 4)
	RunPerft(3, 4)
	RunPerft(4, 4)
	RunPerft(5, 4)
	RunPerft(6, 4)
}

// https://www.chessprogramming.org/Perft_Results#Initial_Position
func RunPerft(position, depth int) {
	fen := getFENPosition(position)

	fmt.Printf("Running perft with depth %d with position %d\n", depth, position)

	depths := []int{}
	for i := 1; i <= depth; i++ {
		depths = append(depths, i)
	}

	for _, d := range depths {
		start := time.Now()
		g := game.NewGameFromFen(fen)
		numPositions := perft(g, d)
		fmt.Printf("Depth: %d, Result: %d, Time: %v", d, numPositions, time.Since(start))
		actual := actualResults[position][d]
		if numPositions != actual {
			fmt.Printf(" - INCORRECT, expected %d\n", actual)
		} else {
			fmt.Println()
		}
	}
}

func perftDivide(g *game.Game, startDepth, currentDepth int) (map[string]uint64, uint64) {
	moves := g.GenerateLegalMoves()

	if currentDepth == 1 {
		return nil, uint64(len(moves))
	}
	results := make(map[string]uint64)
	numLocalNodes := uint64(0)

	for _, move := range moves {
		g.MakeMove(move)
		_, numMovesForThisNode := perftDivide(g, startDepth, currentDepth-1)
		numLocalNodes += numMovesForThisNode
		g.UnmakeMove(move)

		fromFile := move.StartSquare % 8
		fromRank := move.StartSquare / 8
		fromMoveString := string(rune(fromFile+'a')) + string(rune(fromRank+'1'))
		toFile := move.TargetSquare % 8
		toRank := move.TargetSquare / 8
		toMoveString := string(rune(toFile+'a')) + string(rune(toRank+'1'))
		if currentDepth == startDepth {
			results[fromMoveString+toMoveString] += numMovesForThisNode
		}
	}
	return results, numLocalNodes
}

func RunPerftDivide(position, depth int) {
	fen := getFENPosition(position)

	g := game.NewGameFromFen(fen)
	results, numNodes := perftDivide(g, depth, depth)

	keys := make([]string, 0, len(results))
	for key := range results {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, results[key])
	}
	fmt.Printf("Total nodes: %d\n", numNodes)
}

func getFENPosition(position int) string {
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
	} else if position == -1 {
		// debugging position
		fen = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
	}
	return fen
}
