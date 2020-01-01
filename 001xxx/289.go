import "math"

func minFallingPathSum(arr [][]int) int {
	l := len(arr)
	sa, sb, p := 0, 0, -1
	for i := 0; i < l; i++ {
		sa1, sb1, p1 := math.MaxInt64, math.MaxInt64, -1
		for j := 0; j < l; j++ {
			s := arr[i][j]
			if p != j {
				s += sa
			} else {
				s += sb
			}
			if s < sa1 {
				sa1, sb1, p1 = s, sa1, j
			} else if s < sb1 {
				sb1 = s
			}
		}
		sa, sb, p = sa1, sb1, p1
	}
	return sa
}
