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

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func tallestBillboard(rods []int) int {
	m := map[int]int{0: 0}
	oa := (&OrderedArray{}).Init([]int{0},
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b },
	)
	for _, i := range rods {
		l := len(m)
		ps := make([][2]int, l)
		for i := 0; i < l; i++ {
			d := oa.Get(i)
			ps[i] = [2]int{d, m[d]}
		}
		for _, p := range ps {
			d, j := p[0], p[1]
			k := i + d
			v, ok := m[k]
			if !ok {
				oa.Add(k)
			}
			m[k] = max(v, j)
			k = abs(d - i)
			v, ok = m[k]
			if !ok {
				oa.Add(k)
			}
			m[k] = max(v, j+min(d, i))
		}
	}
	return m[oa.Get(0)]
}
