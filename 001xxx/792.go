import "math"
import "math/rand"

type node struct {
	d, b int
	g    float64
}

func (n *node) cal() {
	n.g = float64(n.d) / float64(n.b*(n.b+1))
}

func newNode(d, b int) *node {
	n := &node{d, b, 0.0}
	n.cal()
	return n
}

func (n *node) incB(v int) {
	n.b += v
	n.cal()
}

func (n *node) val(g float64) int {
	return int(math.Sqrt(float64(n.d)*g+0.25)+0.5+1e-8) - n.b
}

func cal(ar []node, r int) {
	if r == 0 {
		return
	}
	n := len(ar)
	g := ar[rand.Intn(n)].g
	i := 0
	for j := n - 1; i <= j; {
		for ; ar[i].g < g; i++ {
		}
		for ; ar[j].g > g; j-- {
		}
		if i <= j {
			ar[i], ar[j] = ar[j], ar[i]
			i, j = i+1, j-1
		}
	}

	g = 1.0 / g
	s := 0
	for j := i; j < n; j++ {
		s += ar[j].val(g)
	}
	if s >= r {
		cal(ar[i:], r)
	} else {
		for j := i; j < n; j++ {
			v := ar[j].val(g)
			ar[j].incB(v)
			r -= v
		}
		if n == 1 {
			ar[0].incB(1)
			r--
		}
		cal(ar, r)
	}
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	ar := make([]node, 0, 16)
	for _, c := range classes {
		if p, t := c[0], c[1]; p != t {
			ar = append(ar, *(newNode(t-p, t)))
		}
	}
	if 0 == len(ar) {
		return 1
	}
	cal(ar, extraStudents)

	o := 0.0
	for i := 0; i < len(ar); i++ {
		o -= float64(ar[i].d) / float64(ar[i].b)
	}
	return 1.0 + o/float64(len(classes))
}
