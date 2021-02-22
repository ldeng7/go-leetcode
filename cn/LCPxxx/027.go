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

func (oa *OrderedArray) RemoveOne(item oaElemType) {
	i := oa.LowerBound(item)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], item) {
		return
	}
	oa.RemoveAt(i)
}

type BlackBox struct {
	pos [][2]int
	neg [][2]int
	m   []map[int]int
	ks  []*OrderedArray
}

func (this *BlackBox) create(n, m, i, dir int) {
	l := len(this.m)
	this.m = append(this.m, map[int]int{})
	this.ks = append(this.ks, (&OrderedArray{}).Init(nil,
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b },
	))
	for j := 0; (dir == 1 && this.pos[i][0] == -1) || (dir == -1 && this.neg[i][0] == -1); j++ {
		if dir == 1 {
			this.pos[i] = [2]int{l, j}
			i = (n+m)*2 - i
		} else {
			this.neg[i] = [2]int{l, j}
			if i <= m*2 {
				i = m*2 - i
			} else {
				i = (m*2+n)*2 - i
			}
		}
		if i != 0 && i != m && i != m+n && i != m*2+n {
			dir = -dir
		}
	}
}

func Constructor(n int, m int) BlackBox {
	l := (n + m) * 2
	this := &BlackBox{
		pos: make([][2]int, l),
		neg: make([][2]int, l),
	}
	for i := 0; i < l; i++ {
		this.pos[i] = [2]int{-1, -1}
		this.neg[i] = [2]int{-1, -1}
	}
	for i := 0; i < l; i++ {
		if i != 0 && i != m+n && this.pos[i][0] == -1 {
			this.create(n, m, i, 1)
		}
		if i != m && i != m*2+n && this.neg[i][0] == -1 {
			this.create(n, m, i, -1)
		}
	}
	return *this
}

func (this *BlackBox) add(p [2]int, index int) {
	i, j := p[0], p[1]
	if _, ok := this.m[i][j]; !ok {
		this.ks[i].Add(j)
	}
	this.m[i][j] = index
}

func (this *BlackBox) Open(index int, dir int) int {
	if p := this.pos[index]; p[0] != -1 {
		this.add(p, index)
	}
	if p := this.neg[index]; p[0] != -1 {
		this.add(p, index)
	}
	var p [2]int
	if dir == 1 {
		p = this.pos[index]
	} else {
		p = this.neg[index]
	}
	i := p[0]
	ks := this.ks[i]
	k := ks.UpperBound(p[1])
	if k != ks.Len() {
		return this.m[i][ks.Get(k)]
	} else {
		return this.m[i][ks.Get(0)]
	}
}

func (this *BlackBox) Close(index int) {
	if p := this.pos[index]; p[0] != -1 {
		i, j := p[0], p[1]
		delete(this.m[i], j)
		this.ks[i].RemoveOne(j)
	}
	if p := this.neg[index]; p[0] != -1 {
		i, j := p[0], p[1]
		delete(this.m[i], j)
		this.ks[i].RemoveOne(j)
	}
}
