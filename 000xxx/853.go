import "sort"

type elemType = float64
type keyType = int

type OrderedArray struct {
	arr    []keyType
	lessCb func(a, b keyType) bool
}

func (oa *OrderedArray) Init(arr []keyType, lessCb func(keyType, keyType) bool) *OrderedArray {
	oa.arr = arr
	if len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool { return lessCb(arr[i], arr[j]) })
	}
	oa.lessCb = lessCb
	return oa
}

func (oa *OrderedArray) Len() int {
	return len(oa.arr)
}

func (oa *OrderedArray) Get(index int) keyType {
	return oa.arr[index]
}

func (oa *OrderedArray) BinSearch(item keyType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := int(uint(i+j) >> 1)
		if oa.lessCb(oa.arr[h], item) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) Index(item keyType) int {
	i := oa.BinSearch(item)
	if i == len(oa.arr) || oa.arr[i] != item {
		return -1
	}
	return i
}

func (oa *OrderedArray) Count(item keyType) int {
	i := oa.BinSearch(item)
	if i == len(oa.arr) || oa.arr[i] != item {
		return 0
	}
	ie := oa.BinSearch(item + 1)
	return ie - i
}

func (oa *OrderedArray) Add(item keyType) {
	i := oa.BinSearch(item)
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

func (oa *OrderedArray) Remove(item keyType) bool {
	i := oa.BinSearch(item)
	if i == len(oa.arr) || oa.arr[i] != item {
		return false
	}
	oa.RemoveAt(i)
	return true
}

func (oa *OrderedArray) RemoveRange(indexBegin, indexEnd int) {
	if indexEnd != len(oa.arr) {
		copy(oa.arr[indexBegin:], oa.arr[indexEnd:])
	}
	oa.arr = oa.arr[:len(oa.arr)-(indexEnd-indexBegin)]
}

func (oa *OrderedArray) RemoveEach(item keyType) int {
	i := oa.BinSearch(item)
	if i == len(oa.arr) || oa.arr[i] != item {
		return 0
	}
	ie := oa.BinSearch(item + 1)
	oa.RemoveRange(i, ie)
	return ie - i
}

type ArrayOrderedMap struct {
	m  map[keyType]elemType
	oa OrderedArray
}

func (om *ArrayOrderedMap) Init(lessCb func(keyType, keyType) bool) *ArrayOrderedMap {
	om.m = map[keyType]elemType{}
	om.oa.Init(nil, lessCb)
	return om
}

func (om *ArrayOrderedMap) Len() int {
	return len(om.m)
}

func (om *ArrayOrderedMap) Get(key keyType) (elemType, bool) {
	v, ok := om.m[key]
	return v, ok
}

func (om *ArrayOrderedMap) GetAt(index int) (keyType, elemType) {
	k := om.oa.Get(index)
	v, _ := om.m[k]
	return k, v
}

func (om *ArrayOrderedMap) BinSearch(key keyType) int {
	return om.oa.BinSearch(key)
}

func (om *ArrayOrderedMap) Set(key keyType, value elemType) {
	if _, ok := om.m[key]; !ok {
		om.oa.Add(key)
	}
	om.m[key] = value
}

func (om *ArrayOrderedMap) Remove(key keyType) {
	delete(om.m, key)
	om.oa.Remove(key)
}

func (om *ArrayOrderedMap) RemoveRange(indexBegin, indexEnd int) {
	for i := indexBegin; i < indexEnd; i++ {
		delete(om.m, om.oa.Get(i))
	}
	om.oa.RemoveRange(indexBegin, indexEnd)
}

func carFleet(target int, position []int, speed []int) int {
	o := 0
	var t0 float64
	m := (&ArrayOrderedMap{}).Init(func(a, b keyType) bool { return a < b })
	for i, p := range position {
		m.Set(-p, float64(target-p)/float64(speed[i]))
	}
	l := m.Len()
	for i := 0; i < l; i++ {
		_, t := m.GetAt(i)
		if t > t0 {
			t0, o = t, o+1
		}
	}
	return o
}
