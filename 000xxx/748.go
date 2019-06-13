import "sort"

func shortestCompletingWord(licensePlate string, words []string) string {
	m := map[int][][]byte{}
	for _, w := range words {
		m[len(w)] = append(m[len(w)], []byte(w))
	}
	ls := []int{}
	for i, _ := range m {
		ls = append(ls, i)
	}
	sort.Ints(ls)
	bs := []byte{}
	for _, b := range []byte(licensePlate) {
		if b >= 'a' && b <= 'z' {
			bs = append(bs, b)
		} else if b >= 'A' && b <= 'Z' {
			bs = append(bs, b+'a'-'A')
		}
	}

	for _, i := range ls {
		if i < len(bs) {
			continue
		}
		for _, s := range m[i] {
			ok := true
			c := map[byte]int{}
			for _, b := range s {
				c[b] = c[b] + 1
			}
			for _, b := range bs {
				c[b]--
				if c[b] < 0 {
					ok = false
					break
				}
			}
			if ok {
				return string(s)
			}
		}
	}
	return ""
}
