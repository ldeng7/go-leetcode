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
	m  map[int]float64
	oa OrderedArray
}

func (om *OrderedMap) Init(lessCb func(int, int) bool) *OrderedMap {
	om.m = map[int]float64{}
	om.oa.Init(nil, lessCb)
	return om
}

func (om *OrderedMap) Range(f func(k int, v float64) bool) int {
	i := 0
	for _, k := range om.oa.Arr {
		if !f(k, om.m[k]) {
			break
		}
		i++
	}
	return i
}

func (om *OrderedMap) Set(k int, v float64) {
	if _, ok := om.m[k]; !ok {
		om.oa.Add(k)
	}
	om.m[k] = v
}

func carFleet(target int, position []int, speed []int) int {
	o := 0
	var t0 float64
	m := (&OrderedMap{}).Init(nil)
	for i, p := range position {
		m.Set(-p, float64(target-p)/float64(speed[i]))
	}
	m.Range(func(_ int, t float64) bool {
		if t > t0 {
			t0, o = t, o+1
		}
		return true
	})
	return o
}
