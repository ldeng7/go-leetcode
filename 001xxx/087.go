import "sort"

type Node struct {
	c    byte
	cs   []byte
	l, w int
}

func parse(S string) []Node {
	ls, i, ns := len(S), 0, []Node{}
	for i < ls {
		if S[i] != '{' {
			i, ns = i+1, append(ns, Node{c: S[i], l: 1})
			continue
		}
		if S[i+1] == '}' {
			i += 2
			continue
		}
		cs := []byte{}
		for {
			cs = append(cs, S[i+1])
			i += 2
			if S[i] == '}' {
				break
			}
		}
		if len(cs) != 1 {
			sort.Slice(cs, func(i, j int) bool { return cs[i] < cs[j] })
			i, ns = i+1, append(ns, Node{cs: cs, l: len(cs)})
		} else {
			i, ns = i+1, append(ns, Node{c: cs[0], l: 1})
		}
	}
	return ns
}

func expand(S string) []string {
	ns := parse(S)
	l, w := len(ns), 1
	for i := l - 1; i >= 0; i-- {
		ns[i].w, w = w, w*ns[i].l
	}

	o := make([]string, w)
	for i := 0; i < w; i++ {
		bs := make([]byte, l)
		for j := 0; j < l; j++ {
			n := &ns[j]
			if n.l == 1 {
				bs[j] = n.c
			} else {
				bs[j] = n.cs[i/n.w%n.l]
			}
		}
		o[i] = string(bs)
	}
	return o
}
