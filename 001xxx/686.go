import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	s, l := 0, len(aliceValues)
	for i, b := range bobValues {
		aliceValues[i] += b
		s -= b
	}
	sort.Ints(aliceValues)
	for i := l - 1; i >= 0; i -= 2 {
		s += aliceValues[i]
	}
	if s > 0 {
		return 1
	} else if s < 0 {
		return -1
	}
	return 0
}
