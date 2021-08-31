import "sort"

func getMaximumConsecutive(coins []int) int {
	o := 1
	sort.Ints(coins)
	for _, c := range coins {
		if c > o {
			return o
		}
		o += c
	}
	return o
}
