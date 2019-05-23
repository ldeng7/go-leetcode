import (
	"fmt"
	"sort"
)

func largestTimeFromDigits(A []int) string {
	sort.Ints(A)
	makeMap := func() map[int]int {
		m := map[int]int{}
		for _, a := range A {
			if _, ok := m[a]; !ok {
				m[a] = 1
			} else {
				m[a]++
			}
		}
		return m
	}
	m := makeMap()

	useDigit := func(a int) {
		cnt, _ := m[a]
		if cnt <= 1 {
			delete(m, a)
		} else {
			m[a]--
		}
	}

	resolve := func(a, b int) (string, bool) {
		var c, d int
		for c, _ = range m {
			useDigit(c)
			break
		}
		for d, _ = range m {
		}
		if (d > c && d < 6) || (c > d && c >= 6 && d < 6) {
			c, d = d, c
		}
		if c >= 6 {
			return "", false
		}
		return fmt.Sprintf("%d%d:%d%d", a, b, c, d), true
	}

	if _, ok := m[2]; ok {
		useDigit(2)
		for _, b := range [4]int{3, 2, 1, 0} {
			if _, ok := m[b]; ok {
				useDigit(b)
				if out, ok := resolve(2, b); ok {
					return out
				} else {
					break
				}
			}
		}
	}

	for _, a := range [2]int{1, 0} {
		m = makeMap()
		if _, ok := m[a]; ok {
			useDigit(a)
			b := A[3]
			if b == a {
				b = A[2]
			}
			useDigit(b)
			if out, ok := resolve(a, b); ok {
				return out
			}
		}
	}
	return ""
}
