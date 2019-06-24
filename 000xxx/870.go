import "sort"

type PriorityQueue struct {
	arr    [][2]int
	lessCb func(a, b [2]int) bool
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

func (pq *PriorityQueue) Init(arr [][2]int, lessCb func([2]int, [2]int) bool) *PriorityQueue {
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

func (pq *PriorityQueue) Push(item [2]int) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Pop() ([2]int, bool) {
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

func advantageCount(A []int, B []int) []int {
	n := len(A)
	l, r := 0, n-1
	sort.Ints(A)
	out := make([]int, n)
	q := (&PriorityQueue{}).Init(nil, func(a, b [2]int) bool {
		return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	for i, b := range B {
		q.Push([2]int{b, i})
	}
	for 0 != q.Len() {
		t, _ := q.Pop()
		if A[r] > t[0] {
			out[t[1]], r = A[r], r-1
		} else {
			out[t[1]], l = A[l], l+1
		}
	}
	return out
}
