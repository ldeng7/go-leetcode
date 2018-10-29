import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return false
	}
	iRow := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	})
	iRow--
	if -1 == iRow {
		return false
	}
	row := matrix[iRow]
	iCol := sort.Search(len(row), func(i int) bool {
		return row[i] >= target
	})
	if iCol == len(row) || row[iCol] != target {
		return false
	}
	return true
}
