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

func (oa *OrderedArray) Len() int {
	return len(oa.arr)
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

type MyCalendarThree struct {
	m  map[int]int
	oa OrderedArray
}

func Constructor() MyCalendarThree {
	this := MyCalendarThree{m: map[int]int{}}
	this.oa.Init(nil,
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b },
	)
	return this
}

func (this *MyCalendarThree) Book(start int, end int) int {
	if _, ok := this.m[start]; !ok {
		this.oa.Add(start)
	}
	this.m[start]++
	if _, ok := this.m[end]; !ok {
		this.oa.Add(end)
	}
	this.m[end]--

	c, o, l := 0, 0, this.oa.Len()
	for i := 0; i < l; i++ {
		c += this.m[this.oa.Get(i)]
		if c > o {
			o = c
		}
	}
	return o
}
