import "sort"

type PriorityQueue struct {
	Arr [][2]int
}

func less(a, b [2]int) bool {
	return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !less(pq.Arr[j], pq.Arr[i]) {
			break
		}
		pq.Arr[i], pq.Arr[j] = pq.Arr[j], pq.Arr[i]
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
		if j2 := j1 + 1; j2 < n && less(pq.Arr[j2], pq.Arr[j1]) {
			j = j2
		}
		if !less(pq.Arr[j], pq.Arr[i]) {
			break
		}
		pq.Arr[i], pq.Arr[j] = pq.Arr[j], pq.Arr[i]
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue) Init(arr [][2]int) *PriorityQueue {
	pq.Arr = arr
	l := len(pq.Arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

func (pq *PriorityQueue) Push(item [2]int) {
	pq.Arr = append(pq.Arr, item)
	pq.up(len(pq.Arr) - 1)
}

func (pq *PriorityQueue) Pop() ([2]int, bool) {
	i := len(pq.Arr) - 1
	if i >= 0 {
		pq.Arr[0], pq.Arr[i] = pq.Arr[i], pq.Arr[0]
		pq.down(0, i)
		v := pq.Arr[i]
		pq.Arr = pq.Arr[:i]
		return v, true
	}
	return [2]int{}, false
}

func advantageCount(A []int, B []int) []int {
	n := len(A)
	l, r := 0, n-1
	sort.Ints(A)
	out := make([]int, n)
	q := (&PriorityQueue{}).Init(nil)
	for i, b := range B {
		q.Push([2]int{b, i})
	}
	for 0 != len(q.Arr) {
		t, _ := q.Pop()
		if A[r] > t[0] {
			out[t[1]], r = A[r], r-1
		} else {
			out[t[1]], l = A[l], l+1
		}
	}
	return out
}
