import "sort"

func calFreq(s string) int {
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for _, c := range m {
		if c > 0 {
			return c
		}
	}
	return -1
}

func numSmallerByFrequency(queries []string, words []string) []int {
	l := len(words)
	fs := make([]int, l)
	for i, s := range words {
		fs[i] = calFreq(s)
	}
	sort.Ints(fs)
	o := make([]int, len(queries))
	for i, s := range queries {
		f := calFreq(s)
		o[i] = l - sort.Search(l, func(j int) bool { return fs[j] > f })
	}
	return o
}
