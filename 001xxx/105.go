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

func minHeightShelves(books [][]int, shelfWidth int) int {
	l := len(books)
	t := make([]int, l+1)
	for i := 0; i < l; i++ {
		t[i] = math.MaxInt64
	}
	for i := l - 1; i >= 0; i-- {
		m, lw := 0, shelfWidth
		for j := i; j < l && lw >= books[j][0]; j++ {
			m = max(m, books[j][1])
			t[i] = min(t[i], m+t[j+1])
			lw -= books[j][0]
		}
	}
	return t[0]
}
