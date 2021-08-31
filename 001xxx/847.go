import "sort"

type oaValType = int
type oaValCmpCb = func(oaValType, oaValType) bool

type OrderedArray struct {
	arr    []oaValType
	lessCb oaValCmpCb
	eqCb   oaValCmpCb
}

func (oa *OrderedArray) Init(arr []oaValType, sorted bool, lessCb oaValCmpCb) *OrderedArray {
	oa.arr = arr
	if (!sorted) && len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool {
			return lessCb(arr[i], arr[j])
		})
	}
	oa.lessCb = lessCb
	oa.eqCb = func(rooms, b oaValType) bool {
		return (!lessCb(rooms, b)) && (!lessCb(b, rooms))
	}
	return oa
}

func (oa *OrderedArray) Len() int {
	return len(oa.arr)
}

func (oa *OrderedArray) Get(index int) oaValType {
	return oa.arr[index]
}

func (oa *OrderedArray) LowerBound(val oaValType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := i + (j-i)>>1
		if oa.lessCb(oa.arr[h], val) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) Add(val oaValType) {
	i := oa.LowerBound(val)
	if i != len(oa.arr) {
		var e oaValType
		oa.arr = append(oa.arr, e)
		copy(oa.arr[i+1:], oa.arr[i:])
		oa.arr[i] = val
	} else {
		oa.arr = append(oa.arr, val)
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func closestRoom(rooms [][]int, queries [][]int) []int {
	sort.Slice(rooms, func(i, j int) bool {
		a, b := rooms[i], rooms[j]
		return a[1] > b[1] || (a[1] == b[1] && a[0] > b[0])
	})

	qs := make([][3]int, len(queries))
	for i, q := range queries {
		qs[i][0], qs[i][1], qs[i][2] = q[1], q[0], i
	}
	sort.Slice(qs, func(i, j int) bool {
		a, b := qs[i], qs[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
	})

	s := (&OrderedArray{}).Init(make([]int, 0, 32), true,
		func(a, b int) bool { return a < b })
	i := 0
	o := make([]int, len(queries))
	for _, t := range qs {
		for ; i < len(rooms) && rooms[i][1] >= t[0]; i++ {
			s.Add(rooms[i][0])
		}
		if 0 == s.Len() {
			o[t[2]] = -1
			continue
		}
		idp := t[1]
		if j := s.LowerBound(idp); j != s.Len() {
			id := s.Get(j)
			if j != 0 {
				if id1 := s.Get(j - 1); abs(id1-idp) <= abs(id-idp) {
					id = id1
				}
			}
			o[t[2]] = id
		} else {
			o[t[2]] = s.Get(s.Len() - 1)
		}
	}
	return o
}
