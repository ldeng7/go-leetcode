type PriorityQueue struct {
	arr    []int
	lessCb func(a, b int) bool
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

func (pq *PriorityQueue) Init(arr []int, lessCb func(int, int) bool) *PriorityQueue {
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

func (pq *PriorityQueue) Top() (int, bool) {
	if len(pq.arr) != 0 {
		return pq.arr[0], true
	}
	return 0, false
}

func (pq *PriorityQueue) Push(item int) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Set(index, item int) {
	pq.arr[index] = item
	if !pq.down(index, len(pq.arr)) {
		pq.up(index)
	}
}

type KthLargest struct {
	k int
	q PriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
	this := KthLargest{k: k}
	if k > len(nums) {
		k = len(nums)
	}
	(&this.q).Init(nums[:k], func(a, b int) bool { return a < b })
	for _, n := range nums[k:] {
		if t, _ := this.q.Top(); n > t {
			this.q.Set(0, n)
		}
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	if this.q.Len() == this.k {
		if t, _ := this.q.Top(); val > t {
			this.q.Set(0, val)
		}
	} else {
		this.q.Push(val)
	}
	t, _ := this.q.Top()
	return t
}
