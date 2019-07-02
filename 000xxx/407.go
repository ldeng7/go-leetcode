import "math"

type pqElemType = [3]int
type pqElemLessCb = func(pqElemType, pqElemType) bool

type PriorityQueue struct {
	arr    []pqElemType
	lessCb pqElemLessCb
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.lessCb(pq.arr[j], pq.arr[i]) {
			break
		}
		pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
		j = i
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

func (pq *PriorityQueue) Init(arr []pqElemType, lessCb pqElemLessCb) *PriorityQueue {
	pq.arr = arr
	pq.lessCb = lessCb
	l := len(pq.arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

func (pq *PriorityQueue) Len() int {
	return len(pq.arr)
}

func (pq *PriorityQueue) Push(item pqElemType) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Pop() *pqElemType {
	i := len(pq.arr) - 1
	if i >= 0 {
		pq.arr[0], pq.arr[i] = pq.arr[i], pq.arr[0]
		pq.down(0, i)
		e := pq.arr[i]
		pq.arr = pq.arr[:i]
		return &e
	}
	return nil
}

func (pq *PriorityQueue) Abandon() []pqElemType {
	return pq.arr
}

var dirs = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func trapRainWater(heightMap [][]int) int {
	if 0 == len(heightMap) || 0 == len(heightMap[0]) {
		return 0
	}
	o, mx, m, n := 0, math.MinInt64, len(heightMap), len(heightMap[0])
	q := (&PriorityQueue{}).Init(nil, func(a, b [3]int) bool { return a[0] < b[0] })
	t := make([][]bool, m)
	ie, je := m-1, n-1
	for i := 0; i < m; i++ {
		t[i] = make([]bool, n)
		q.Push([3]int{heightMap[i][0], i, 0})
		q.Push([3]int{heightMap[i][je], i, je})
		t[i][0], t[i][je] = true, true
	}
	for j := 1; j < n-1; j++ {
		q.Push([3]int{heightMap[0][j], 0, j})
		q.Push([3]int{heightMap[ie][j], ie, j})
		t[0][j], t[ie][j] = true, true
	}

	for 0 != q.Len() {
		p := *q.Pop()
		if p[0] > mx {
			mx = p[0]
		}
		for _, d := range dirs {
			y, x := p[1]+d[0], p[2]+d[1]
			if y < 0 || y >= m || x < 0 || x >= n || t[y][x] {
				continue
			}
			t[y][x] = true
			h := heightMap[y][x]
			if h < mx {
				o += mx - h
			}
			q.Push([3]int{h, y, x})
		}
	}
	return o
}
