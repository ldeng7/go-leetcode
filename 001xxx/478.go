import (
	"math"
	"sort"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minDistance(houses []int, K int) int {
	sort.Ints(houses)
	l := len(houses)
	ar := make([][]int, l)
	for i := 0; i < l; i++ {
		ar[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			m := i + (j-i)>>1
			for k := i; k <= j; k++ {
				ar[i][j] += abs(houses[k] - houses[m])
			}
		}
	}

	t := make([][]int, l)
	for i := 0; i < l; i++ {
		tt := make([]int, K+1)
		for j := 0; j <= K; j++ {
			tt[j] = math.MaxInt32
		}
		tt[1] = ar[0][i]
		t[i] = tt
	}
	for i := 0; i < l; i++ {
		for j := 2; j <= min(i+1, K); j++ {
			for k := j - 1; k <= i; k++ {
				t[i][j] = min(t[i][j], t[k-1][j-1]+ar[k][i])
			}
		}
	}
	return t[l-1][K]
}
