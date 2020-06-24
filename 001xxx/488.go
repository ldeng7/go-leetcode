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

func (oa *OrderedArray) UpperBound(item oaElemType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := i + (j-i)>>1
		if !oa.lessCb(item, oa.arr[h]) {
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

func avoidFlood(rains []int) []int {
	l, n := len(rains), 0
	o := make([]int, l)
	for i, r := range rains {
		o[i] = -1
		if r == 0 {
			n++
		}
	}
	m := map[int]int{}
	oa := (&OrderedArray{}).Init(make([]int, 0, n),
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b },
	)

	for i, r := range rains {
		if r == 0 {
			oa.Add(i)
		} else if j, ok := m[r]; !ok {
			m[r] = i
		} else {
			if oa.Len() == 0 {
				return nil
			}
			k := oa.UpperBound(j)
			if k == oa.Len() {
				return nil
			}
			o[oa.Get(k)], m[r] = r, i
			oa.RemoveAt(k)
		}
	}
	for i := 0; i < oa.Len(); i++ {
		o[oa.Get(i)] = 1
	}
	return o
}
