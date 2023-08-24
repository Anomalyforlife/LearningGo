package chessboard

type File []bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	count := 0
	for _, isOccupied := range cb[file] {
		if isOccupied {
			count++
		}
	}
	return count
}
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	count := 0
	for file := range cb {
		if cb[file][rank-1] {
			count++
		}
	}
	return count
}
func CountAll(cb Chessboard) int {
	return 64
}
func CountOccupied(cb Chessboard) int {
	count := 0
	for file := range cb {
		count += CountInFile(cb, file)
	}
	return count
}
