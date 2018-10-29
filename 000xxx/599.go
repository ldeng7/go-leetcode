import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	out := []string{}
	m := map[string]int{}
	n := math.MaxInt64
	for i, s := range list1 {
		m[s] = i
	}
	for i, s := range list2 {
		if _, ok := m[s]; !ok {
			continue
		}
		sum := i + m[s]
		if sum == n {
			out = append(out, s)
		} else if sum < n {
			n = sum
			out = []string{s}
		}
	}
	return out
}
