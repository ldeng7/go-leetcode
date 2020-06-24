import "sort"

type pqElemType = int
type pqElemCmpCb = func(pqElemType, pqElemType) bool

type PriorityQueue struct {
	arr    []pqElemType
	lessCb pqElemCmpCb
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

func (pq *PriorityQueue) Init(arr []pqElemType, lessCb pqElemCmpCb) *PriorityQueue {
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

func (pq *PriorityQueue) Top() *pqElemType {
	if len(pq.arr) != 0 {
		e := pq.arr[0]
		return &e
	}
	return nil
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	pairs := make([]uint64, n)
	for i := 0; i < n; i++ {
		pairs[i] = (uint64(efficiency[i]) << 32) | uint64(speed[i])
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i] < pairs[j] })
	pq := (&PriorityQueue{}).Init([]int{}, func(a, b int) bool { return a < b })
	o, s := 0, 0
	for i := n - 1; i >= 0; i-- {
		pair := pairs[i]
		e, sp := int(pair>>32), int(pair&0xffffffff)
		s += sp
		pq.Push(sp)
		if pq.Len() > k {
			t := pq.Pop()
			s -= *t
		}
		o = max(o, s*e)
	}
	return o % 1000000007
}
