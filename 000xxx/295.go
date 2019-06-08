type HeapArray struct {
	Arr   []int
	cmpCb func(a, b int) bool
}

func (ha *HeapArray) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !ha.cmpCb(ha.Arr[j], ha.Arr[i]) {
			break
		}
		ha.Arr[i], ha.Arr[j] = ha.Arr[j], ha.Arr[i]
		j = i
	}
}

func (ha *HeapArray) down(i0, n int) bool {
	i := i0
	for {
		j1 := i<<1 + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && ha.cmpCb(ha.Arr[j2], ha.Arr[j1]) {
			j = j2
		}
		if !ha.cmpCb(ha.Arr[j], ha.Arr[i]) {
			break
		}
		ha.Arr[i], ha.Arr[j] = ha.Arr[j], ha.Arr[i]
		i = j
	}
	return i > i0
}

func (ha *HeapArray) init() {
	l := len(ha.Arr)
	for i := l>>1 - 1; i >= 0; i-- {
		ha.down(i, l)
	}
}

func NewHeapArray(arr []int) *HeapArray {
	ha := &HeapArray{
		Arr:   arr,
		cmpCb: func(a, b int) bool { return a < b },
	}
	ha.init()
	return ha
}

func (ha *HeapArray) Push(item int) {
	ha.Arr = append(ha.Arr, item)
	ha.up(len(ha.Arr) - 1)
}

func (ha *HeapArray) Pop() (int, bool) {
	i := len(ha.Arr) - 1
	if i >= 0 {
		ha.Arr[0], ha.Arr[i] = ha.Arr[i], ha.Arr[0]
		ha.down(0, i)
		v := ha.Arr[i]
		ha.Arr = ha.Arr[:i]
		return v, true
	}
	return 0, false
}

type MedianFinder struct {
	s, l *HeapArray
}

func Constructor() MedianFinder {
	return MedianFinder{NewHeapArray(nil), NewHeapArray(nil)}
}

func (this *MedianFinder) AddNum(num int) {
	this.s.Push(num)
	this.l.Push(-this.s.Arr[0])
	this.s.Pop()
	if len(this.s.Arr) < len(this.l.Arr) {
		this.s.Push(-this.l.Arr[0])
		this.l.Pop()
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.s.Arr) > len(this.l.Arr) {
		return float64(this.s.Arr[0])
	}
	return (float64(this.s.Arr[0]) - float64(this.l.Arr[0])) / 2.0
}
