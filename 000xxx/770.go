import (
	"sort"
	"strconv"
	"strings"
)

type Term struct {
	i    int
	strs []string
}

type Terms []Term

func (ts Terms) Len() int      { return len(ts) }
func (ts Terms) Swap(i, j int) { ts[i], ts[j] = ts[j], ts[i] }
func (ts Terms) Less(i, j int) bool {
	a, b := ts[i], ts[j]
	la, lb := len(a.strs), len(b.strs)
	if la != lb {
		return la > lb
	}
	for i := 0; i < la; i++ {
		sa, sb := a.strs[i], b.strs[i]
		if sa != sb {
			return sa < sb
		}
	}
	return false
}

func atoi(expr string, i int) (int, int) {
	v, c, ie := 0, expr[i], len(expr)
	for {
		v, i = v*10+int(c-'0'), i+1
		if i == ie {
			break
		}
		c = expr[i]
		if c < '0' || c > '9' {
			break
		}
	}
	return v, i
}

func parseVar(expr string, i int) int {
	j, ie := i+1, len(expr)
	for ; j != ie; j++ {
		c := expr[j]
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			break
		}
	}
	return j
}

func mergeStrs(a, b []string) []string {
	la, lb := len(a), len(b)
	o := make([]string, la+lb)
	for i, j, k := 0, 0, 0; i < la || j < lb; k++ {
		if i < la && j < lb {
			if a[i] < b[j] {
				o[k], i = a[i], i+1
			} else {
				o[k], j = b[j], j+1
			}
		} else {
			if i < la {
				o[k], i = a[i], i+1
			} else {
				o[k], j = b[j], j+1
			}
		}
	}
	return o
}

func eqStrs(a, b []string) bool {
	la, lb := len(a), len(b)
	if la != lb {
		return false
	}
	for i, sa := range a {
		if sa != b[i] {
			return false
		}
	}
	return true
}

func cal(expr string, i int, m map[string]int) ([]Term, int) {
	ar, lop := Terms{}, Terms{}
	ie, b, neg := len(expr), false, false
	for i != ie && expr[i] != ')' {
		c := expr[i]
		if c == ' ' || c == '*' {
			i++
			continue
		} else if c == '+' || c == '-' {
			for _, term := range lop {
				ar = append(ar, term)
			}
			lop, b = []Term{}, false
			if c == '-' {
				neg = true
			}
			i++
			continue
		}

		rop := Terms{}
		if c >= '0' && c <= '9' {
			v, j := atoi(expr, i)
			rop, i = append(rop, Term{v, []string{}}), j
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			j := parseVar(expr, i)
			s := expr[i:j]
			if v, ok := m[s]; ok {
				rop = append(rop, Term{v, []string{}})
			} else {
				rop = append(rop, Term{1, []string{s}})
			}
			i = j
		} else if c == '(' {
			rop, i = cal(expr, i+1, m)
		}

		if b {
			terms := Terms{}
			for _, tl := range lop {
				for _, tr := range rop {
					v := tl.i * tr.i
					if neg {
						v = -v
					}
					terms = append(terms, Term{v, mergeStrs(tl.strs, tr.strs)})
				}
			}
			sort.Sort(terms)
			lop = terms
		} else {
			if neg {
				for j := 0; j < len(rop); j++ {
					rop[j].i *= -1
				}
			}
			lop = rop
			b = true
		}
		neg = false
	}

	if b {
		for _, term := range lop {
			ar = append(ar, term)
		}
	}
	sort.Sort(ar)
	o := []Term{}
	for _, term := range ar {
		e := len(o) - 1
		if 0 == len(o) || !eqStrs(term.strs, o[e].strs) {
			o = append(o, term)
		} else if o[e].i == -term.i {
			o = o[:e]
		} else {
			o[e].i += term.i
		}
	}
	if i != ie {
		i++
	}
	return o, i
}

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	o, m := []string{}, map[string]int{}
	for i, e := range evalvars {
		m[e] = evalints[i]
	}

	terms, _ := cal(expression, 0, m)
	for _, term := range terms {
		if term.i != 0 {
			if len(term.strs) != 0 {
				o = append(o, strconv.Itoa(term.i)+"*"+strings.Join(term.strs, "*"))
			} else {
				o = append(o, strconv.Itoa(term.i))
			}
		}
	}
	return o
}
