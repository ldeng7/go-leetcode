import "sort"

func maxCoins(piles []int) int {
	sort.Ints(piles)
	o, l := 0, len(piles)
	for i, j := 0, l-1; i < l/3; i, j = i+1, j-2 {
		o += piles[j-1]
	}
	return o
}
