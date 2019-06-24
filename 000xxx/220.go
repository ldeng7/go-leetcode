import "sort"

type OrderedArray struct {
	arr    []int
	lessCb func(a, b int) bool
}

func (oa *OrderedArray) Init(arr []int, lessCb func(int, int) bool) *OrderedArray {
	oa.arr = arr
	if len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool { return lessCb(arr[i], arr[j]) })
	}
	oa.lessCb = lessCb
	return oa
}

func (oa *OrderedArray) Len() int {
	return len(oa.arr)
}

func (oa *OrderedArray) Get(index int) int {
	return oa.arr[index]
}

func (oa *OrderedArray) BinSearch(item int) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := int(uint(i+j) >> 1)
		if oa.lessCb(oa.arr[h], item) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) Add(item int) {
	i := oa.BinSearch(item)
	if i != len(oa.arr) {
		oa.arr = append(oa.arr, 0)
		copy(oa.arr[i+1:], oa.arr[i:])
		oa.arr[i] = item
	} else {
		oa.arr = append(oa.arr, item)
	}
}

func (oa *OrderedArray) RemoveAt(index int) {
	if index != len(oa.arr)-1 {
		copy(oa.arr[index:], oa.arr[index+1:])
	}
	oa.arr = oa.arr[:len(oa.arr)-1]
}

func (oa *OrderedArray) Remove(item int) bool {
	i := oa.BinSearch(item)
	if i == len(oa.arr) || oa.arr[i] != item {
		return false
	}
	oa.RemoveAt(i)
	return true
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	oa := (&OrderedArray{}).Init(nil, func(a, b int) bool { return a < b })
	j := 0
	for i, num := range nums {
		if i-j > k {
			oa.Remove(nums[j])
			j++
		}
		h := oa.BinSearch(num - t)
		if h != oa.Len() && oa.Get(h) <= num+t {
			return true
		}
		oa.Add(num)
	}
	return false
}
