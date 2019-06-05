import "sort"

type AutocompleteSystem struct {
	m  map[string]int
	bs []byte
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	this := &AutocompleteSystem{map[string]int{}, []byte{}}
	for i, s := range sentences {
		this.m[s] += times[i]
	}
	return *this
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.m[string(this.bs)]++
		this.bs = []byte{}
		return []string{}
	}

	this.bs = append(this.bs, c)
	q := []string{}
	for s, f := range this.m {
		i := 0
		for ; i < len(this.bs) && i < len(s); i++ {
			if this.bs[i] != s[i] {
				break
			}
		}
		if i == len(this.bs) {
			j := sort.Search(len(q), func(k int) bool {
				s1 := q[k]
				f1 := this.m[s1]
				if f == f1 {
					return s > s1
				}
				return f < f1
			})
			if j != len(q) {
				q = append(q, "")
				copy(q[j+1:], q[j:])
				q[j] = s
			} else {
				q = append(q, s)
			}
			if len(q) > 3 {
				q = q[1:]
			}
		}
	}

	out := make([]string, len(q))
	for i, s := range q {
		out[len(q)-i-1] = s
	}
	return out
}
