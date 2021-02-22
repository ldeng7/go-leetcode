const MOD = 1000000007

type BitTree struct {
	ar []int
	n  int
}

func NewBitTree(n int) *BitTree {
	return &BitTree{make([]int, n+1), n}
}

func (bt *BitTree) Add(p int, v int) {
	p++
	for p <= bt.n {
		bt.ar[p] += v
		p += p & -p
	}
}

func (bt *BitTree) Sum(p int) int {
	o := 0
	p++
	for p > 0 {
		o += bt.ar[p]
		p -= p & -p
	}
	return o
}

func bonus(n int, leadership [][]int, operations [][]int) []int {
	subs := make([][]int, n)
	for _, l := range leadership {
		a := l[0] - 1
		subs[a] = append(subs[a], l[1]-1)
	}

	arIn := make([]int, n)
	arOut := make([]int, n)
	i := 0
	var dfs func(int)
	dfs = func(a int) {
		arIn[a], i = i, i+1
		for _, b := range subs[a] {
			dfs(b)
		}
		arOut[a] = i - 1
	}
	dfs(0)

	bt1, bt2 := NewBitTree(n), NewBitTree(n)
	add := func(l, r, v int) {
		bt1.Add(l, v)
		bt1.Add(r+1, -v)
		bt2.Add(l, v*(l-1))
		bt2.Add(r+1, -v*r)
	}
	sum := func(l, r int) int {
		l--
		o := ((bt1.Sum(r)*r - bt2.Sum(r)) - (bt1.Sum(l)*l - bt2.Sum(l))) % MOD
		if o < 0 {
			o += MOD
		}
		return o
	}

	o := make([]int, 0, 32)
	for _, op := range operations {
		a := op[1] - 1
		switch op[0] {
		case 1:
			add(arIn[a], arIn[a], op[2])
		case 2:
			add(arIn[a], arOut[a], op[2])
		case 3:
			o = append(o, sum(arIn[a], arOut[a]))
		}
	}
	return o
}
