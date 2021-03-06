type pqElemType = string
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

func findItinerary(tickets [][]string) []string {
	o, st := []string{}, []string{"JFK"}
	m := map[string]*PriorityQueue{}
	for _, t := range tickets {
		k := t[0]
		pq := m[k]
		if nil == pq {
			m[k] = (&PriorityQueue{}).Init(nil, func(a, b string) bool { return a < b })
			pq = m[k]
		}
		pq.Push(t[1])
	}
	for 0 != len(st) {
		s := st[len(st)-1]
		pq := m[s]
		if nil != pq && 0 != pq.Len() {
			st = append(st, *pq.Pop())
		} else {
			o = append(o, s)
			st = st[:len(st)-1]
		}
	}
	l := len(o)
	for i := 0; i < len(o)>>1; i++ {
		o[i], o[l-i-1] = o[l-i-1], o[i]
	}
	return o
}
