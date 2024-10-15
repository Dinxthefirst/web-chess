package game

// func checkLine(fromRow, fromCol, toRow, toCol int, g *Game) bool {
// 	rowDir := getDirection(toRow - fromRow)
// 	colDir := getDirection(toCol - fromCol)
// 	row, col := fromRow+rowDir, fromCol+colDir
// 	for row != toRow || col != toCol {
// 		if g.Board[row][col] != nil {
// 			return false
// 		}
// 		row += rowDir
// 		col += colDir
// 	}
// 	return true
// }

// func getDirection(delta int) int {
// 	if delta == 0 {
// 		return 0
// 	}
// 	if delta > 0 {
// 		return 1
// 	}
// 	return -1
// }

// func checkLinear(fromRow, fromCol, toRow, toCol int, g *Game) bool {
// 	if fromRow == toRow || fromCol == toCol {
// 		return checkLine(fromRow, fromCol, toRow, toCol, g)
// 	}
// 	return false
// }

// func checkDiagonal(fromRow, fromCol, toRow, toCol int, g *Game) bool {
// 	if abs(fromRow-toRow) == abs(fromCol-toCol) {
// 		return checkLine(fromRow, fromCol, toRow, toCol, g)
// 	}
// 	return false
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func checkKing(fromRow, fromCol, toRow, toCol int, g *Game) bool {
// 	if abs(fromRow-toRow) <= 1 && abs(fromCol-toCol) <= 1 {
// 		if g.Board[toRow][toCol] == nil || g.Board[toRow][toCol].Color() != g.ActiveColor {
// 			return true
// 		}
// 	}
// 	return false
// }

// func checkKnight(fromRow, fromCol, toRow, toCol int, g *Game) bool {
// 	if !(abs(fromRow-toRow) == 1 && abs(fromCol-toCol) == 2) || (abs(fromRow-toRow) == 2 && abs(fromCol-toCol) == 1) {
// 		return false
// 	}
// 	if g.Board[toRow][toCol] != nil {
// 		return false
// 	}
// 	return true
// }
