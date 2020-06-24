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

func (oa *OrderedArray) RemoveAt(index int) {
	if index != len(oa.arr)-1 {
		copy(oa.arr[index:], oa.arr[index+1:])
	}
	oa.arr = oa.arr[:len(oa.arr)-1]
}

func (oa *OrderedArray) RemoveOne(item oaElemType) {
	i := oa.LowerBound(item)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], item) {
		return
	}
	oa.RemoveAt(i)
}

type ExamRoom struct {
	n  int
	m  map[int]bool
	oa OrderedArray
}

func Constructor(N int) ExamRoom {
	this := ExamRoom{n: N, m: map[int]bool{}}
	this.oa.Init(nil,
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b },
	)
	return this
}

func (this *ExamRoom) Seat() int {
	b, m, k := 0, 0, 0
	l := len(this.m)
	for i := 0; i < l; i++ {
		j := this.oa.Get(i)
		if b == 0 {
			if m < j-b {
				m, k = j-b, 0
			}
		} else {
			if t := (j - b + 1) >> 1; m < t {
				m, k = t, b+t-1
			}
		}
		b = j + 1
	}
	if b > 0 && m < this.n-b {
		k = this.n - 1
	}
	if _, ok := this.m[k]; !ok {
		this.oa.Add(k)
	}
	this.m[k] = true
	return k
}

func (this *ExamRoom) Leave(p int) {
	this.oa.RemoveOne(p)
	delete(this.m, p)
}
