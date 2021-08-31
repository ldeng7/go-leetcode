type pqValType = uint64
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

func countRestrictedPaths(n int, edges [][]int) int {
	m := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		m[u] = append(m[u], [2]int{v, w})
		m[v] = append(m[v], [2]int{u, w})
	}
	t, ds := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		ds[i] = 1e10
	}
	t[n-1], ds[n-1] = 1, 0
	pq := (&PriorityQueue{}).Init(make([]uint64, 0, 64), false, func(a, b uint64) bool {
		return a < b
	})
	pq.Push(uint64(n - 1))

	for 0 != pq.Len() {
		p := *(pq.Pop())
		w, u := int(p>>32), int(p&0xffffffff)
		if w != ds[u] {
			continue
		}
		for _, p := range m[u] {
			v, d := p[0], p[1]
			if s := w + d; s < ds[v] {
				ds[v] = s
				pq.Push(uint64(s<<32) | uint64(v))
			} else if w > ds[v] {
				t[u] = (t[u] + t[v]) % 1000000007
			}
		}
		if u == 0 {
			break
		}
	}
	return t[0]
}
