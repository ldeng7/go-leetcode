import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	t := make([]int, n+1)
	for i := 0; i <= n; i++ {
		t[i] = math.MaxInt64
	}
	t[n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			t[j] = max(1, min(t[j], t[j+1])-dungeon[i][j])
		}
	}
	return t[0]
}
