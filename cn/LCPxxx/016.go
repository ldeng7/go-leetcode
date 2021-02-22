func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Triangle struct {
	a, b, c       int
	va, vb, vc, v int
}

func NewTriangle(a, b, c, va, vb, vc int) *Triangle {
	return &Triangle{a, b, c, va, vb, vc, va + vb + vc}
}

func (t *Triangle) plus(t1 *Triangle, cv int) int {
	if nil == t || t.v == 0 {
		return t1.v
	} else if t1.v == 0 {
		return t.v
	}
	a, b, c, a1, b1, c1 := t.a, t.b, t.c, t1.a, t1.b, t1.c
	if (a == a1 && b == b1) || (a == b1 && b == c1) || (a == a1 && b == c1) {
		return t.vc + t1.v
	} else if (b == a1 && c == b1) || (b == b1 && c == c1) || (b == a1 && c == c1) {
		return t.va + t1.v
	} else if (a == a1 && c == b1) || (a == b1 && c == c1) || (a == a1 && c == c1) {
		return t.vb + t1.v
	}
	return t.v + t1.v - cv
}

type Vertex struct {
	Sum, V     int
	t0, t1, t2 *Triangle
}

func (v *Vertex) Add(t *Triangle) {
	t0, t1, t2 := v.t0, v.t1, v.t2
	v.Sum = max(max(v.Sum, t0.plus(t, v.V)), max(t1.plus(t, v.V), t2.plus(t, v.V)))
	if nil == t0 || t.v >= t0.v {
		v.t0, v.t1, v.t2 = t, t0, t1
	} else if nil == t1 || t.v >= t1.v {
		v.t1, v.t2 = t, t1
	} else if nil == t2 || t.v >= t2.v {
		v.t2 = t
	}
}

func maxWeight(edges [][]int, value []int) int {
	l := len(value)
	edgeSet := map[int]bool{}
	edgeCounts := make([]int, l)
	for _, e := range edges {
		a, b := e[0], e[1]
		if a > b {
			a, b = b, a
			e[0], e[1] = a, b
		}
		edgeSet[(a<<16)|b] = true
		edgeCounts[a]++
	}
	edgePoints := make([]int, l+1)
	for i := 1; i < l; i++ {
		edgePoints[i] = edgePoints[i-1] + edgeCounts[i-1]
	}

	edgeTo := make([]int, len(edges))
	for _, e := range edges {
		a := e[0]
		edgeTo[edgePoints[a]] = e[1]
		edgePoints[a]++
	}
	edgePoints[0] = 0
	for i := 1; i < l+1; i++ {
		edgePoints[i] = edgePoints[i-1] + edgeCounts[i-1]
	}

	vs := make([]Vertex, l)
	for i, v := range value {
		vs[i].V = v
	}
	for a := 0; a < l; a++ {
		for i := edgePoints[a]; i < edgePoints[a+1]; i++ {
			b := edgeTo[i]
			ma, mi := a, b
			if edgeCounts[a] < edgeCounts[b] {
				ma, mi = b, a
			}
			for j := edgePoints[mi]; j < edgePoints[mi+1]; j++ {
				c := edgeTo[j]
				if ma <= c && edgeSet[(ma<<16)|c] {
					t := NewTriangle(a, b, c, value[a], value[b], value[c])
					vs[a].Add(t)
					vs[b].Add(t)
					vs[c].Add(t)
				}
			}
		}
	}

	o := 0
	for i := 0; i < l; i++ {
		o = max(o, vs[i].Sum)
	}
	return o
}
