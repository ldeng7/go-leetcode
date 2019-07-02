import "sort"

func canReorderDoubled(A []int) bool {
	m := map[int]int{}
	for _, a := range A {
		m[a]++
	}
	sort.Ints(A)
	for _, a := range A {
		if 0 == m[a] {
			continue
		}
		var k int
		if a < 0 {
			k = a / 2
		} else {
			k = a * 2
		}
		if 0 == m[k] {
			return false
		} else {
			m[k]--
		}
		m[a]--
	}
	return true
}
