import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func suggestedProducts(products []string, searchWord string) [][]string {
	l, ls := len(products), len(searchWord)
	o := make([][]string, ls)
	for i := 0; i < ls; i++ {
		j := 0
		for k := 0; k < l; k++ {
			if p := products[k]; i < len(p) && searchWord[i] == p[i] {
				products[j] = p
				j++
			}
		}
		sort.Strings(products[:j])
		la := min(j, 3)
		ar := make([]string, la)
		copy(ar, products[:la])
		l = j
		o[i] = ar
	}
	return o
}
