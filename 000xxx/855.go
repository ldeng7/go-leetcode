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

func (oa *OrderedArray) Index(item int) int {
	i := oa.binSearch(item)
	if i == len(oa.Arr) || oa.Arr[i] != item {
		return -1
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

func (oa *OrderedArray) RemoveAt(index int) {
	if index != len(oa.Arr)-1 {
		copy(oa.Arr[index:], oa.Arr[index+1:])
	}
	oa.Arr = oa.Arr[:len(oa.Arr)-1]
}

type OrderedMap struct {
	m  map[int]bool
	oa OrderedArray
}

func (om *OrderedMap) Init(lessCb func(int, int) bool) *OrderedMap {
	om.m = map[int]bool{}
	om.oa.Init(nil, lessCb)
	return om
}

func (om *OrderedMap) Range(f func(k int, v bool) bool) int {
	i := 0
	for _, k := range om.oa.Arr {
		if !f(k, om.m[k]) {
			break
		}
		i++
	}
	return i
}

func (om *OrderedMap) Set(k int, v bool) {
	if _, ok := om.m[k]; !ok {
		om.oa.Add(k)
	}
	om.m[k] = v
}

func (om *OrderedMap) Remove(k int) {
	delete(om.m, k)
	om.oa.RemoveAt(om.oa.Index(k))
}

type ExamRoom struct {
	n int
	m *OrderedMap
}

func Constructor(N int) ExamRoom {
	return ExamRoom{N, (&OrderedMap{}).Init(nil)}
}

func (this *ExamRoom) Seat() int {
	b, m, i := 0, 0, 0
	this.m.Range(func(j int, _ bool) bool {
		if b == 0 {
			if m < j-b {
				m, i = j-b, 0
			}
		} else {
			t := (j - b + 1) >> 1
			if m < t {
				m, i = t, b+t-1
			}
		}
		b = j + 1
		return true
	})
	if b > 0 && m < this.n-b {
		m, i = this.n-b, this.n-1
	}
	this.m.Set(i, true)
	return i
}

func (this *ExamRoom) Leave(p int) {
	this.m.Remove(p)
}
