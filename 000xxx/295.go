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

func (pq *PriorityQueue) Pop() (int, bool) {
	i := len(pq.arr) - 1
	if i >= 0 {
		pq.arr[0], pq.arr[i] = pq.arr[i], pq.arr[0]
		pq.down(0, i)
		v := pq.arr[i]
		pq.arr = pq.arr[:i]
		return v, true
	}
	return 0, false
}

type MedianFinder struct {
	s, l *PriorityQueue
}

func Constructor() MedianFinder {
	return MedianFinder{
		(&PriorityQueue{}).Init(nil, func(a, b int) bool { return a < b }),
		(&PriorityQueue{}).Init(nil, func(a, b int) bool { return a < b }),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.s.Push(num)
	ts, _ := this.s.Top()
	this.l.Push(-ts)
	this.s.Pop()
	if this.s.Len() < this.l.Len() {
		tl, _ := this.l.Top()
		this.s.Push(-tl)
		this.l.Pop()
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.s.Len() > this.l.Len() {
		t, _ := this.s.Top()
		return float64(t)
	}
	ts, _ := this.s.Top()
	tl, _ := this.l.Top()
	return (float64(ts) - float64(tl)) / 2.0
}
