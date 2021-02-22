type pqValType = int
type pqValCmpCb = func(pqValType, pqValType) bool

type PriorityQueue struct {
	arr    []pqValType
	lessCb pqValCmpCb
}

func (pq *PriorityQueue) up(i int) {
	for {
		j := (i - 1) / 2
		if j == i || !pq.lessCb(pq.arr[i], pq.arr[j]) {
			break
		}
		pq.arr[j], pq.arr[i] = pq.arr[i], pq.arr[j]
		i = j
	}
}

func (pq *PriorityQueue) down(i0, n int) bool {
	i := i0
	for {
		j1 := i<<1 + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && pq.lessCb(pq.arr[j2], pq.arr[j1]) {
			j = j2
		}
		if !pq.lessCb(pq.arr[j], pq.arr[i]) {
			break
		}
		pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue) Init(arr []pqValType, sorted bool, lessCb pqValCmpCb) *PriorityQueue {
	pq.arr = arr
	pq.lessCb = lessCb
	if !sorted {
		l := len(pq.arr)
		for i := l>>1 - 1; i >= 0; i-- {
			pq.down(i, l)
		}
	}
	return pq
}

func (pq *PriorityQueue) Top() *pqValType {
	if len(pq.arr) != 0 {
		e := pq.arr[0]
		return &e
	}
	return nil
}

func (pq *PriorityQueue) Push(val pqValType) {
	pq.arr = append(pq.arr, val)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Set(index int, val pqValType) {
	pq.arr[index] = val
	if !pq.down(index, len(pq.arr)) {
		pq.up(index)
	}
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	i, l := 0, len(heights)-1
	pq := (&PriorityQueue{}).Init(make([]int, 0, 16), true, func(a, b int) bool { return a > b })
	for ; i < l; i++ {
		d := heights[i+1] - heights[i]
		if d <= 0 {
			continue
		}
		if d <= bricks {
			bricks -= d
			pq.Push(d)
		} else if ladders > 0 {
			if t := pq.Top(); nil != t && d < *t {
				bricks += *t - d
				pq.Set(0, d)
			}
			ladders--
		} else {
			break
		}
	}
	return i
}
