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

func (ha *HeapArray) Set(index, item int) {
	ha.Arr[index] = item
	if !ha.down(index, len(ha.Arr)) {
		ha.up(index)
	}
}

func findKthLargest(nums []int, k int) int {
	h := NewHeapArray(nums[:k])
	for _, num := range nums[k:] {
		if num <= h.Arr[0] {
			continue
		}
		h.Set(0, num)
	}
	return h.Arr[0]
}
