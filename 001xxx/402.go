import "sort"

func maxSatisfaction(satisfaction []int) int {
	l := len(satisfaction)
	if 0 == l {
		return 0
	}
	sort.Ints(satisfaction)
	if satisfaction[l-1] <= 0 {
		return 0
	}
	o, t := 0, 0
	for i := l - 1; i >= 0; i-- {
		if t = t + satisfaction[i]; t > 0 {
			o += t
		}
	}
	return o
}
