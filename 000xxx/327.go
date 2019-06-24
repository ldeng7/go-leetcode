import "sort"

type elemType = int
type OrderedArray struct {
	arr    []elemType
	lessCb func(a, b elemType) bool
}

func (oa *OrderedArray) Init(arr []elemType, lessCb func(elemType, elemType) bool) *OrderedArray {
	oa.arr = arr
	if len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool { return lessCb(arr[i], arr[j]) })
	}
	oa.lessCb = lessCb
	return oa
}

func (oa *OrderedArray) BinSearch(item elemType) int {
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

func (oa *OrderedArray) Add(item elemType) {
	i := oa.BinSearch(item)
	if i != len(oa.arr) {
		oa.arr = append(oa.arr, 0)
		copy(oa.arr[i+1:], oa.arr[i:])
		oa.arr[i] = item
	} else {
		oa.arr = append(oa.arr, item)
	}
}

func countRangeSum(nums []int, lower int, upper int) int {
	o, s := 0, 0
	ss := (&OrderedArray{}).Init([]int{0}, func(a, b int) bool { return a < b })
	for _, n := range nums {
		s += n
		o += ss.BinSearch(s-lower+1) - ss.BinSearch(s-upper)
		ss.Add(s)
	}
	return o
}
