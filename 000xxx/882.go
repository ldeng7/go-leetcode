type pqElemType = [2]int
type pqElemLessCb = func(pqElemType, pqElemType) bool

type PriorityQueue struct {
	arr    []pqElemType
	lessCb pqElemLessCb
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.lessCb(pq.arr[j], pq.arr[i]) {
			break
		}
		pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
		j = i
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

func (pq *PriorityQueue) Init(arr []pqElemType, lessCb pqElemLessCb) *PriorityQueue {
	pq.arr = arr
	pq.lessCb = lessCb
	l := len(pq.arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

func (pq *PriorityQueue) Len() int {
	return len(pq.arr)
}

func (pq *PriorityQueue) Push(item pqElemType) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Pop() *pqElemType {
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reachableNodes(edges [][]int, M int, N int) int {
	t := make([][][2]int, N)
	for _, e := range edges {
		i, j := e[0], e[1]
		t[i] = append(t[i], [2]int{j, e[2]})
		t[j] = append(t[j], [2]int{i, e[2]})
	}
	m := map[int]int{0: 0}
	for i := 1; i < N; i++ {
		m[i] = M + 1
	}
	v := map[int]int{}
	q := (&PriorityQueue{}).Init(nil, func(a, b [2]int) bool {
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	q.Push([2]int{0, 0})

	o := 0
	for 0 != q.Len() {
		p := *q.Pop()
		i, j := p[0], p[1]
		if i > m[j] {
			continue
		}
		m[j], o = i, o+1
		for _, p := range t[j] {
			k, w := p[0], p[1]
			v[(j<<16)|k] = min(w, M-i)
			i1 := i + w + 1
			if i1 < min(m[k], M+1) {
				q.Push([2]int{i1, k})
				m[k] = i1
			}
		}
	}
	for _, e := range edges {
		i, j := e[0], e[1]
		o += min(e[2], v[(i<<16)|j]+v[(j<<16)|i])
	}
	return o
}
