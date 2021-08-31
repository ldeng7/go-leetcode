import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	o, s := 0, 0
	for _, c := range costs {
		if t := s + c; t <= coins {
			o, s = o+1, t
		} else {
			break
		}
	}
	return o
}
