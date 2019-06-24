import "sort"

type elemType = int
type keyType = int

type OrderedArray struct {
	arr    []elemType
	lessCb func(a, b elemType) bool
}

func (oa *OrderedArray) Init(arr []elemType, lessCb func(elemType, elemType) bool) *OrderedArray {
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

func (oa *OrderedArray) Get(index int) elemType {
	return oa.arr[index]
}

func (oa *OrderedArray) BinSearch(item elemType) int {
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

func (oa *OrderedArray) Index(item elemType) int {
	i := oa.BinSearch(item)
	if i == len(oa.arr) || oa.arr[i] != item {
		return -1
	}
	return i
}

func (oa *OrderedArray) Count(item elemType) int {
	i := oa.BinSearch(item)
	if i == len(oa.arr) || oa.arr[i] != item {
		return 0
	}
	ie := oa.BinSearch(item + 1)
	return ie - i
}

func (oa *OrderedArray) Add(item elemType) {
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

func (oa *OrderedArray) Remove(item elemType) bool {
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

func (oa *OrderedArray) RemoveEach(item elemType) int {
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

type MyCalendarThree struct {
	om ArrayOrderedMap
}

func Constructor() MyCalendarThree {
	this := MyCalendarThree{}
	this.om.Init(func(a, b keyType) bool { return a < b })
	return this
}

func (this *MyCalendarThree) Book(start int, end int) int {
	v, _ := this.om.Get(start)
	this.om.Set(start, v+1)
	v, _ = this.om.Get(end)
	this.om.Set(end, v-1)
	c, o, l := 0, 0, this.om.Len()
	for i := 0; i < l; i++ {
		_, v := this.om.GetAt(i)
		c += v
		if c > o {
			o = c
		}
	}
	return o
}
