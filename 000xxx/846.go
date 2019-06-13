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

func (oa *OrderedArray) binSearch(item int) int {
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
	i := oa.binSearch(item)
	if i != len(oa.Arr) {
		oa.Arr = append(oa.Arr, 0)
		copy(oa.Arr[i+1:], oa.Arr[i:])
		oa.Arr[i] = item
	} else {
		oa.Arr = append(oa.Arr, item)
	}
}

type OrderedMap struct {
	m  map[int]int
	oa OrderedArray
}

func (om *OrderedMap) Init() *OrderedMap {
	om.m = map[int]int{}
	om.oa.Init(nil, nil)
	return om
}

func (om *OrderedMap) Len() int {
	return len(om.m)
}

func (om *OrderedMap) Get(k int) (int, bool) {
	v, ok := om.m[k]
	return v, ok
}

func (om *OrderedMap) Range(f func(k int, v int) bool) int {
	i := 0
	for _, k := range om.oa.Arr {
		if !f(k, om.m[k]) {
			break
		}
		i++
	}
	return i
}
func (om *OrderedMap) Set(k, v int) {
	if _, ok := om.m[k]; !ok {
		om.oa.Add(k)
	}
	om.m[k] = v
}

func (om *OrderedMap) Inc(k int) {
	if _, ok := om.m[k]; !ok {
		om.oa.Add(k)
	}
	om.m[k]++
}

func isNStraightHand(hand []int, W int) bool {
	m := (&OrderedMap{}).Init()
	for _, h := range hand {
		m.Inc(h)
	}
	c := m.Range(func(h, c int) bool {
		if c == 0 {
			return true
		}
		for i := h; i < h+W; i++ {
			c1, _ := m.Get(i)
			if c1 < c {
				return false
			}
			m.Set(i, c1-c)
		}
		return true
	})
	return c == m.Len()
}
