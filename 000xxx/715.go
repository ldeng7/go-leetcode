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

func (oa *OrderedArray) RemoveRange(indexBegin, indexEnd int) {
	if indexEnd != len(oa.arr) {
		copy(oa.arr[indexBegin:], oa.arr[indexEnd:])
	}
	oa.arr = oa.arr[:len(oa.arr)-(indexEnd-indexBegin)]
}

type RangeModule struct {
	m  map[int]int
	oa OrderedArray
}

func Constructor() RangeModule {
	this := RangeModule{m: map[int]int{}}
	this.oa.Init(nil,
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b },
	)
	return this
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
	l, r := this.oa.LowerBound(left+1), this.oa.LowerBound(right+1)
	if l != 0 {
		l--
		if this.m[this.oa.Get(l)] < left {
			l++
		}
	}
	if l == r {
		return left, right
	}
	left = min(left, this.oa.Get(l))
	right = max(right, this.m[this.oa.Get(r-1)])
	for i := l; i < r; i++ {
		delete(this.m, this.oa.Get(i))
	}
	this.oa.RemoveRange(l, r)
	return left, right
}

func (this *RangeModule) AddRange(left int, right int) {
	l, r := this.find(left, right)
	if _, ok := this.m[l]; !ok {
		this.oa.Add(l)
	}
	this.m[l] = r
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	i := this.oa.LowerBound(left + 1)
	return 0 != i && this.m[this.oa.Get(i-1)] >= right
}

func (this *RangeModule) RemoveRange(left int, right int) {
	l, r := this.find(left, right)
	if left > l {
		if _, ok := this.m[l]; !ok {
			this.oa.Add(l)
		}
		this.m[l] = left
	}
	if right < r {
		if _, ok := this.m[right]; !ok {
			this.oa.Add(right)
		}
		this.m[right] = r
	}
}
