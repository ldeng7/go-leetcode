import "sort"

type CharCount struct {
	b byte
	c int
}

func insertToQueue(q []CharCount, b byte, c int) []CharCount {
	i := sort.Search(len(q), func(j int) bool {
		return q[j].c < c || (q[j].c == c && q[j].b <= b)
	})
	if i == len(q) {
		q = append(q, CharCount{b, c})
	} else {
		q = append(q, CharCount{})
		copy(q[i+1:], q[i:])
		q[i] = CharCount{b, c}
	}
	return q
}

func rearrangeString(s string, k int) string {
	if k == 0 {
		return s
	}
	l, bs, m, q := len(s), []byte{}, map[byte]int{}, []CharCount{}
	for _, b := range []byte(s) {
		m[b]++
	}
	for b, c := range m {
		q = insertToQueue(q, b, c)
	}
	for 0 != len(q) {
		v := []CharCount{}
		cnt := k
		if l < k {
			cnt = l
		}
		for i := 0; i < cnt; i++ {
			if 0 == len(q) {
				return ""
			}
			t := q[0]
			q = q[1:]
			bs = append(bs, t.b)
			t.c, l = t.c-1, l-1
			if t.c > 0 {
				v = append(v, CharCount{t.b, t.c})
			}
		}
		for _, e := range v {
			q = insertToQueue(q, e.b, e.c)
		}
	}
	return string(bs)
}
