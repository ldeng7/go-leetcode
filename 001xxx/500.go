import "sort"

type oaElemType = int
type oaElemCmpCb = func(oaElemType, oaElemType) bool

type OrderedArray struct {
	arr    []oaElemType
	lessCb oaElemCmpCb
	eqCb   oaElemCmpCb
}

func (oa *OrderedArray) Init(arr []oaElemType, sorted bool, lessCb, eqCb oaElemCmpCb) *OrderedArray {
	oa.arr = arr
	if (!sorted) && len(arr) > 1 {
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

func (oa *OrderedArray) Upsert(item oaElemType) {
	i := oa.LowerBound(item)
	if i != len(oa.arr) {
		if !oa.eqCb(oa.arr[i], item) {
			oa.arr = append(oa.arr, 0)
			copy(oa.arr[i+1:], oa.arr[i:])
		}
		oa.arr[i] = item
	} else {
		oa.arr = append(oa.arr, item)
	}
}

func (oa *OrderedArray) RemoveAt(index int) {
	if index != 0 {
		if index != len(oa.arr)-1 {
			copy(oa.arr[index:], oa.arr[index+1:])
		}
		oa.arr = oa.arr[:len(oa.arr)-1]
	} else {
		oa.arr = oa.arr[1:]
	}
}

func (oa *OrderedArray) RemoveOne(item oaElemType) {
	i := oa.LowerBound(item)
	if i == len(oa.arr) || !oa.eqCb(oa.arr[i], item) {
		return
	}
	oa.RemoveAt(i)
}

type FileSharing struct {
	umax int
	uq   *OrderedArray
	cu   []*OrderedArray
	uc   []map[int]bool
}

func Constructor(m int) FileSharing {
	this := &FileSharing{
		umax: 16,
		cu:   make([]*OrderedArray, m+1),
		uc:   make([]map[int]bool, 16+1),
	}

	ar := make([]int, this.umax)
	for i := 0; i < this.umax; i++ {
		ar[i] = i + 1
	}
	this.uq = (&OrderedArray{}).Init(ar, true,
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b })

	for i := 1; i <= m; i++ {
		this.cu[i] = (&OrderedArray{}).Init(make([]int, 0, 16), true,
			func(a, b int) bool { return a < b },
			func(a, b int) bool { return a == b })
	}

	return *this
}

func (this *FileSharing) Join(ownedChunks []int) int {
	var u int
	if 0 != this.uq.Len() {
		u = this.uq.Get(0)
		this.uq.RemoveAt(0)
	} else {
		this.umax++
		u = this.umax
		this.uc = append(this.uc, nil)
	}

	uc := this.uc[u]
	if nil == uc {
		uc = map[int]bool{}
		this.uc[u] = uc
	}
	for _, c := range ownedChunks {
		this.cu[c].Add(u)
		uc[c] = true
	}
	return u
}

func (this *FileSharing) Leave(userID int) {
	this.uq.Add(userID)
	for c, _ := range this.uc[userID] {
		this.cu[c].RemoveOne(userID)
	}
	this.uc[userID] = nil
}

func (this *FileSharing) Request(userID int, chunkID int) []int {
	ar := this.cu[chunkID].arr
	o := make([]int, len(ar))
	if 0 != len(ar) {
		copy(o, ar)
		this.cu[chunkID].Upsert(userID)
		this.uc[userID][chunkID] = true
	}
	return o
}
