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

func busiestServers(k int, arrival []int, load []int) []int {
	m := map[int]int{}
	oa := (&OrderedArray{}).Init(make([]int, 0, len(arrival)), true,
		func(a, b int) bool { return a < b },
		func(a, b int) bool { return a == b },
	)
	ar, cs := make([]int, k), make([]int, k)
	j, ma := -1, 0
	o := make([]int, 0, k)

	cal := func(j int) {
		for {
			h := sort.Search(oa.Len(), func(i int) bool { return oa.Get(i) >= ar[j] })
			if h == oa.Len() {
				break
			}
			cs[j]++
			a := oa.Get(h)
			ar[j] = a + m[a]
			oa.RemoveAt(h)
		}
	}

	for i, a := range arrival {
		j = i % k
		m[a] = load[i]
		oa.arr = append(oa.arr, a)
		cal(j)
	}
	for j, je := (j+1)%k, j; j != je; j = (j + 1) % k {
		cal(j)
	}
	for i, c := range cs {
		if c > ma {
			o = make([]int, 1, k-i)
			o[0] = i
			ma = c
		} else if c == ma {
			o = append(o, i)
		}
	}
	return o
}
