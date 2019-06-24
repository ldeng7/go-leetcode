type elemType = [2]int
type PriorityQueue struct {
	arr    []elemType
	lessCb func(a, b elemType) bool
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

func (pq *PriorityQueue) Init(arr []elemType, lessCb func(elemType, elemType) bool) *PriorityQueue {
	pq.arr = arr
	pq.lessCb = lessCb
	l := len(pq.arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

func (pq *PriorityQueue) Push(item elemType) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Pop() (elemType, bool) {
	i := len(pq.arr) - 1
	if i >= 0 {
		pq.arr[0], pq.arr[i] = pq.arr[i], pq.arr[0]
		pq.down(0, i)
		v := pq.arr[i]
		pq.arr = pq.arr[:i]
		return v, true
	}
	return [2]int{}, false
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	q := (&PriorityQueue{}).Init(nil, func(a, b elemType) bool {
		return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	o := []int{}
	for _, n := range nums {
		m[n]++
	}
	for n, c := range m {
		q.Push([2]int{c, n})
	}
	for i := 0; i < k; i++ {
		p, _ := q.Pop()
		o = append(o, p[1])
	}
	return o
}
