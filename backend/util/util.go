package util

const BoardSize = 8

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FromChessNotation(notation string) int {
	if notation == "-" {
		return -1
	}
	file := int(notation[0] - 'a')
	rank := int(notation[1] - '1')
	return rank*BoardSize + file
}

func ToChessNotation(square int) string {
	if square == -1 {
		return "-"
	}
	file := square % BoardSize
	rank := square / BoardSize
	return string(rune(file+'a')) + string(rune(rank+'1'))
}

// func bitboardString(bitboard uint64) string {
// 	str := ""
// 	for rank := 7; rank >= 0; rank-- {
// 		for file := 0; file < 8; file++ {
// 			str += fmt.Sprintf("%b", bitboard&(1<<(rank*BoardSize+file))>>(rank*BoardSize+file))
// 		}
// 		str += "\n"
// 	}
// 	return str
// }
