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

type Leaderboard struct {
	oa OrderedArray
	m  map[int]int
}

func Constructor() Leaderboard {
	oa := (&OrderedArray{}).Init([]int{},
		func(a, b int) bool { return a > b },
		func(a, b int) bool { return a == b },
	)
	return Leaderboard{*oa, map[int]int{}}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	if s, ok := this.m[playerId]; ok {
		score += s
		this.oa.RemoveOne(s)
	}
	this.m[playerId] = score
	this.oa.Add(score)
}

func (this *Leaderboard) Top(K int) int {
	o := 0
	for i := 0; i < K; i++ {
		o += this.oa.Get(i)
	}
	return o
}

func (this *Leaderboard) Reset(playerId int) {
	this.oa.RemoveOne(this.m[playerId])
	this.oa.Add(0)
	this.m[playerId] = 0
}
