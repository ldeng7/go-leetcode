import "sort"

func heightChecker(heights []int) int {
	t := make([]int, len(heights))
	copy(t, heights)
	sort.Ints(heights)
	c := 0
	for i, h := range heights {
		if h != t[i] {
			c++
		}
	}
	return c
}
