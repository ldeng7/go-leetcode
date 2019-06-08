import "sort"

func numFactoredBinaryTrees(A []int) int {
	o := 0
	m := map[int]int{}
	sort.Ints(A)
	for i, a := range A {
		m[a] = 1
		for j := 0; j < i; j++ {
			b := A[j]
			c := a / b
			if _, ok := m[c]; ok && a%b == 0 {
				m[a] = (m[a] + m[b]*m[c]) % 1000000007
			}
		}
	}
	for _, a := range m {
		o = (o + a) % 1000000007
	}
	return o
}
