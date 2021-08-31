import "sort"

type pqValType = [2]int
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

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	l := len(queries)
	qs := make([][2]int, l)
	for i, q := range queries {
		qs[i][0], qs[i][1] = q, i
	}
	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i], qs[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	pq := (&PriorityQueue{}).Init(make([][2]int, 0, 16), true, func(a, b [2]int) bool {
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	o := make([]int, l)

	for i, j := 0, 0; i < l; i++ {
		q := qs[i]
		v := q[0]
		for ; j < len(intervals); j++ {
			intl := intervals[j]
			if intl[0] > v {
				break
			}
			pq.Push([2]int{intl[1] - intl[0] + 1, intl[1]})
		}
		for 0 != pq.Len() {
			if t := pq.Top(); (*t)[1] >= v {
				break
			}
			pq.Pop()
		}
		if 0 == pq.Len() {
			o[q[1]] = -1
		} else {
			t := pq.Top()
			o[q[1]] = (*t)[0]
		}
	}
	return o
}
