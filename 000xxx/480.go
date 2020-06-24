import "sort"

type oaElemType = int
type oaElemCmpCb = func(oaElemType, oaElemType) bool

type OrderedArray struct {
	arr    []oaElemType
	lessCb oaElemCmpCb
	eqCb   oaElemCmpCb
}

func (oa *OrderedArray) Init(arr []oaElemType, lessCb, eqCb oaElemCmpCb) *OrderedArray {
	oa.arr = arr
	if len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool { return lessCb(arr[i], arr[j]) })
	}
	oa.lessCb = lessCb
	oa.eqCb = eqCb
	return oa
}

func (oa *OrderedArray) Get(index int) oaElemType {
	return oa.arr[index]
}

func (oa *OrderedArray) LowerBound(item oaElemType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := i + (j-i)>>1
		if oa.lessCb(oa.arr[h], item) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) Exist(item oaElemType) (bool, int) {
	i := oa.LowerBound(item)
	return i != len(oa.arr) && !oa.eqCb(oa.arr[i], item), i
}

func (oa *OrderedArray) Add(item oaElemType) {
	i := oa.LowerBound(item)
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

func less(a, b int) bool { return a < b }
func eq(a, b int) bool   { return a == b }
func medianSlidingWindow(nums []int, k int) []float64 {
	l := len(nums)
	o, i := make([]float64, l-k+1), 1
	ar := make([]int, k)
	copy(ar, nums[:k])
	oa := (&OrderedArray{}).Init(ar, less, eq)
	median := func() float64 {
		if k&1 == 1 {
			return float64(oa.Get(k >> 1))
		}
		return float64(oa.Get(k>>1)+oa.Get((k>>1)-1)) / 2
	}

	o[0] = median()
	for j := k; j < l; i, j = i+1, j+1 {
		_, idx := oa.Exist(nums[j-k])
		println(nums[j-k], nums[j], idx, i)
		oa.RemoveAt(idx)
		oa.Add(nums[j])
		o[i] = median()
	}
	return o
}
