type pqValType = int
type pqValCmpCb = func(pqValType, pqValType) bool

type PriorityQueue struct {
	arr    []pqValType
	lessCb pqValCmpCb
}

func (pq *PriorityQueue) up(i int) {
	for {
		j := (i - 1) / 2
		if j == i || !pq.lessCb(pq.arr[i], pq.arr[j]) {
			break
		}
		pq.arr[j], pq.arr[i] = pq.arr[i], pq.arr[j]
		i = j
	}
}

func (pq *PriorityQueue) down(i0, n int) bool {
	i := i0
	for {
		j1 := i<<1 + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && pq.lessCb(pq.arr[j2], pq.arr[j1]) {
			j = j2
		}
		if !pq.lessCb(pq.arr[j], pq.arr[i]) {
			break
		}
		pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue) Init(arr []pqValType, sorted bool, lessCb pqValCmpCb) *PriorityQueue {
	pq.arr = arr
	pq.lessCb = lessCb
	if !sorted {
		l := len(pq.arr)
		for i := l>>1 - 1; i >= 0; i-- {
			pq.down(i, l)
		}
	}
	return pq
}

func (pq *PriorityQueue) Len() int {
	return len(pq.arr)
}

func (pq *PriorityQueue) Push(val pqValType) {
	pq.arr = append(pq.arr, val)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Pop() *pqValType {
	i := len(pq.arr) - 1
	if i >= 0 {
		pq.arr[0], pq.arr[i] = pq.arr[i], pq.arr[0]
		pq.down(0, i)
		e := pq.arr[i]
		pq.arr = pq.arr[:i]
		return &e
	}
	return nil
}

func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	if start == end {
		return 0
	}
	cnt++
	n := len(charge)
	g := make([][][2]int, n)
	for _, p := range paths {
		a, b, d := p[0], p[1], p[2]
		g[a] = append(g[a], [2]int{b, d})
		g[b] = append(g[b], [2]int{a, d})
	}

	tenc := func(a, p int) int {
		return a*cnt + p
	}
	tdec := func(k int) (int, int) {
		return k / cnt, k % cnt
	}
	qenc := func(d, k int) int {
		return (d << 16) | k
	}
	qdec := func(v int) (int, int) {
		return v >> 16, v & 0xffff
	}
	t := make([]bool, n*cnt)
	ds := make([]int, n*cnt)
	for i := 0; i < n*cnt; i++ {
		ds[i] = 100000
	}
	ts := tenc(start, 0)
	ds[ts] = 0
	q := (&PriorityQueue{}).Init(make([]int, 0, 64), true,
		func(a, b int) bool { return a < b })
	q.Push(qenc(0, ts))

	o := 0
	for 0 != q.Len() {
		d, k := qdec(*q.Pop())
		if t[k] {
			continue
		}
		t[k] = true
		a, p := tdec(k)
		if a == end {
			o = d
			break
		}
		for _, pair := range g[a] {
			b, l := pair[0], pair[1]
			if k1 := tenc(b, p-l); p >= l && d+l < ds[k1] {
				ds[k1] = d + l
				q.Push(qenc(d+l, k1))
			}
		}
		if t := d + charge[a]; p < cnt-1 && t < ds[k+1] {
			ds[k+1] = t
			q.Push(qenc(t, k+1))
		}
	}
	return o
}
