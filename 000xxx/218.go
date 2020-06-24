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

func getSkyline(buildings [][]int) [][]int {
	o, ps := [][]int{}, [][2]int{}
	oa := (&OrderedArray{}).Init([]int{0},
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b },
	)
	for _, b := range buildings {
		ps = append(ps, [2]int{b[0], -b[2]}, [2]int{b[1], b[2]})
	}
	sort.Slice(ps, func(i, j int) bool {
		pa, pb := ps[i], ps[j]
		return pa[0] < pb[0] || (pa[0] == pb[0] && pa[1] < pb[1])
	})

	i, j := 0, 0
	for _, p := range ps {
		if p[1] < 0 {
			oa.Add(-p[1])
		} else {
			oa.RemoveOne(p[1])
		}
		i = oa.Get(oa.Len() - 1)
		if i != j {
			o, j = append(o, []int{p[0], i}), i
		}
	}
	return o
}
