import "sort"

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

func enc(t, i int) int {
	return (t << 32) | i
}

func dec(v int) (int, int) {
	return v >> 32, v & 0xffffffff
}

func getOrder(tasks [][]int) []int {
	l := len(tasks)
	ar := make([]int, l)
	for i, t := range tasks {
		ar[i] = enc(t[0], i)
	}
	sort.Ints(ar)
	pq := (&PriorityQueue{}).Init(make([]int, 0, l), true,
		func(a, b int) bool { return a < b })
	i, c := 0, 0
	o := make([]int, 0, l)
	for len(o) < l {
		for ; i < l; i++ {
			t, j := dec(ar[i])
			if t > c {
				break
			}
			pq.Push(enc(tasks[j][1], j))
		}
		if 0 != pq.Len() {
			t, j := dec(*(pq.Pop()))
			c += t
			o = append(o, j)
		} else if i < l {
			c, _ = dec(ar[i])
		}
	}
	return o
}
