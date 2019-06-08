type HeapArray struct {
	Arr   [][2]int
	cmpCb func(a, b [2]int) bool
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

func NewHeapArray(arr [][2]int) *HeapArray {
	ha := &HeapArray{
		Arr:   arr,
		cmpCb: func(a, b [2]int) bool { return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1]) },
	}
	ha.init()
	return ha
}

func (ha *HeapArray) Push(item [2]int) {
	ha.Arr = append(ha.Arr, item)
	ha.up(len(ha.Arr) - 1)
}

func (ha *HeapArray) Pop() ([2]int, bool) {
	i := len(ha.Arr) - 1
	if i >= 0 {
		ha.Arr[0], ha.Arr[i] = ha.Arr[i], ha.Arr[0]
		ha.down(0, i)
		v := ha.Arr[i]
		ha.Arr = ha.Arr[:i]
		return v, true
	}
	return [2]int{}, false
}

func maxSlidingWindow(nums []int, k int) []int {
	out := []int{}
	h := NewHeapArray(nil)
	for i, _ := range nums {
		for 0 != len(h.Arr) && h.Arr[0][1] <= i-k {
			h.Pop()
		}
		h.Push([2]int{nums[i], i})
		if i >= k-1 {
			out = append(out, h.Arr[0][0])
		}
	}
	return out
}
