import "sort"

func check(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}

func palindromePairs(words []string) [][]int {
	o := [][]int{}
	m := map[string]int{}
	s, ss := map[int]bool{}, []int{}
	for i, w := range words {
		m[w] = i
		l := len(w)
		if !s[l] {
			ss = append(ss, l)
			s[l] = true
		}
	}
	sort.Ints(ss)
	for i, w := range words {
		l := len(w)
		bs := []byte(w)
		for j := 0; j < l>>1; j++ {
			bs[j], bs[l-j-1] = bs[l-j-1], bs[j]
		}
		w = string(bs)
		if j, ok := m[w]; ok && j != i {
			o = append(o, []int{i, j})
		}
		for j := 0; j < len(ss) && ss[j] != l; j++ {
			n := ss[j]
			if check(w, 0, l-n-1) {
				if v, ok := m[w[l-n:]]; ok {
					o = append(o, []int{i, v})
				}
			}
			if check(w, n, l-1) {
				if v, ok := m[w[:n]]; ok {
					o = append(o, []int{v, i})
				}
			}
		}
	}
	return o
}
