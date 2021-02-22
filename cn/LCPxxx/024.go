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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func numsGame(nums []int) []int {
	l := len(nums)
	if l == 1 {
		return []int{0}
	}
	for i := 0; i < l; i++ {
		nums[i] -= i
	}
	s0, s1 := min(nums[0], nums[1]), max(nums[0], nums[1])
	q0 := (&PriorityQueue{}).Init([]int{s0}, func(a, b int) bool { return a > b })
	q1 := (&PriorityQueue{}).Init([]int{s1}, func(a, b int) bool { return a < b })
	o := make([]int, l)
	o[1] = s1 - s0

	for i := 2; i < l; i++ {
		if n := nums[i]; n <= *(q0.Top()) {
			q0.Push(n)
			s0 += n
		} else {
			q1.Push(n)
			s1 += n
		}
		if d := q0.Len() - q1.Len(); d == 2 {
			u := *(q0.Pop())
			q1.Push(u)
			s0, s1 = s0-u, s1+u
		} else if d == -1 {
			u := *(q1.Pop())
			q0.Push(u)
			s0, s1 = s0+u, s1-u
		}
		d := s1 - s0
		if i&1 == 0 {
			d += *(q0.Top())
		}
		o[i] = d % 1000000007
	}
	return o
}

