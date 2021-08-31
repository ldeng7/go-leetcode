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

func (pq *PriorityQueue) Top() *pqValType {
	if len(pq.arr) != 0 {
		e := pq.arr[0]
		return &e
	}
	return nil
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

func (pq *PriorityQueue) Set(index int, val pqValType) {
	pq.arr[index] = val
	if !pq.down(index, len(pq.arr)) {
		pq.up(index)
	}
}

const MOD = 1000000007

func enc(p, a int) uint64 {
	return (uint64(p) << 32) | uint64(a)
}

func dec(v uint64) (int, int) {
	return int(v >> 32), int(v & 0xffffffff)
}

func getNumberOfBacklogOrders(orders [][]int) int {
	pq1 := (&PriorityQueue{}).Init(make([]uint64, 0, 64), true,
		func(a, b uint64) bool { return a < b })
	pq2 := (&PriorityQueue{}).Init(make([]uint64, 0, 64), true,
		func(a, b uint64) bool { return a > b })
	for _, od := range orders {
		p, a := od[0], od[1]
		if od[2] == 0 {
			for a > 0 && 0 != pq1.Len() {
				p1, a1 := dec(*(pq1.Top()))
				if p1 > p {
					break
				}
				if a1 > a {
					pq1.Set(0, enc(p1, a1-a))
					a = 0
				} else {
					pq1.Pop()
					a -= a1
				}
			}
			if a > 0 {
				pq2.Push(enc(p, a))
			}
		} else {
			for a > 0 && 0 != pq2.Len() {
				p1, a1 := dec(*(pq2.Top()))
				if p1 < p {
					break
				}
				if a1 > a {
					pq2.Set(0, enc(p1, a1-a))
					a = 0
				} else {
					pq2.Pop()
					a -= a1
				}
			}
			if a > 0 {
				pq1.Push(enc(p, a))
			}
		}
	}

	o := 0
	for pq1.Len() != 0 {
		_, a := dec(*(pq1.Pop()))
		o = (o + a) % MOD
	}
	for pq2.Len() != 0 {
		_, a := dec(*(pq2.Pop()))
		o = (o + a) % MOD
	}
	return o
}
