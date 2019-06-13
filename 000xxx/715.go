import "sort"

type OrderedArray struct {
	Arr    []int
	lessCb func(a, b int) bool
}

func (oa *OrderedArray) Len() int           { return len(oa.Arr) }
func (oa *OrderedArray) Less(i, j int) bool { return oa.lessCb(oa.Arr[i], oa.Arr[j]) }
func (oa *OrderedArray) Swap(i, j int)      { oa.Arr[i], oa.Arr[j] = oa.Arr[j], oa.Arr[i] }

func (oa *OrderedArray) Init(arr []int, lessCb func(int, int) bool) *OrderedArray {
	oa.Arr = arr
	if len(arr) > 1 {
		sort.Sort(oa)
	}
	oa.lessCb = lessCb
	if nil == lessCb {
		oa.lessCb = func(a, b int) bool { return a < b }
	}
	return oa
}

func (oa *OrderedArray) BinSearch(item int) int {
	i, j := 0, len(oa.Arr)
	for i < j {
		h := int(uint(i+j) >> 1)
		if oa.lessCb(oa.Arr[h], item) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) Add(item int) {
	i := oa.BinSearch(item)
	if i != len(oa.Arr) {
		oa.Arr = append(oa.Arr, 0)
		copy(oa.Arr[i+1:], oa.Arr[i:])
		oa.Arr[i] = item
	} else {
		oa.Arr = append(oa.Arr, item)
	}
}

func (oa *OrderedArray) RemoveRange(indexBegin, indexEnd int) {
	if indexEnd != len(oa.Arr) {
		copy(oa.Arr[indexBegin:], oa.Arr[indexEnd:])
	}
	oa.Arr = oa.Arr[:len(oa.Arr)-(indexEnd-indexBegin)]
}

type OrderedMap struct {
	m  map[int]int
	oa OrderedArray
}

func (om *OrderedMap) Init(lessCb func(int, int) bool) *OrderedMap {
	om.m = map[int]int{}
	om.oa.Init(nil, lessCb)
	return om
}

func (om *OrderedMap) Len() int {
	return len(om.m)
}

func (om *OrderedMap) GetAt(index int) int {
	v, _ := om.m[om.oa.Arr[index]]
	return v
}

func (om *OrderedMap) KeyAt(index int) int {
	return om.oa.Arr[index]
}

func (om *OrderedMap) KeyBinSearch(k int) int {
	return om.oa.BinSearch(k)
}

func (om *OrderedMap) Set(k int, v int) {
	if _, ok := om.m[k]; !ok {
		om.oa.Add(k)
	}
	om.m[k] = v
}

func (om *OrderedMap) RemoveRange(indexBegin, indexEnd int) {
	for i := indexBegin; i < indexEnd; i++ {
		delete(om.m, om.oa.Arr[i])
	}
	om.oa.RemoveRange(indexBegin, indexEnd)
}

type RangeModule struct {
	om OrderedMap
}

func Constructor() RangeModule {
	this := &RangeModule{OrderedMap{}}
	this.om.Init(nil)
	return *this
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func (this *RangeModule) find(left, right int) (int, int) {
	l, r := this.om.KeyBinSearch(left+1), this.om.KeyBinSearch(right+1)
	if l != 0 {
		l--
		if this.om.GetAt(l) < left {
			l++
		}
	}
	if l == r {
		return left, right
	}
	left, right = min(left, this.om.KeyAt(l)), max(right, this.om.GetAt(r-1))
	this.om.RemoveRange(l, r)
	return left, right
}

func (this *RangeModule) AddRange(left int, right int) {
	l, r := this.find(left, right)
	this.om.Set(l, r)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	i := this.om.KeyBinSearch(left + 1)
	return i != 0 && this.om.GetAt(i-1) >= right
}

func (this *RangeModule) RemoveRange(left int, right int) {
	l, r := this.find(left, right)
	if left > l {
		this.om.Set(l, left)
	}
	if right < r {
		this.om.Set(right, r)
	}
}
