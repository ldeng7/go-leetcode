import (
	"math"
	"sort"
)

type PriorityQueue struct {
	arr    []float64
	lessCb func(a, b float64) bool
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

func (pq *PriorityQueue) Init(arr []float64, lessCb func(float64, float64) bool) *PriorityQueue {
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

func (pq *PriorityQueue) Push(item float64) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Pop() (float64, bool) {
	i := len(pq.arr) - 1
	if i >= 0 {
		pq.arr[0], pq.arr[i] = pq.arr[i], pq.arr[0]
		pq.down(0, i)
		v := pq.arr[i]
		pq.arr = pq.arr[:i]
		return v, true
	}
	return 0, false
}

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	ps := make([][2]float64, len(quality))
	for i, q := range quality {
		ps[i] = [2]float64{float64(wage[i]) / float64(q), float64(q)}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i][0] < ps[j][0]
	})

	var out, s float64 = math.MaxFloat64, 0
	q := (&PriorityQueue{}).Init(nil, func(a, b float64) bool { return a < b })
	for _, p := range ps {
		s += p[1]
		q.Push(-p[1])
		if q.Len() > K {
			qu, _ := q.Pop()
			s += qu
		}
		if q.Len() == K {
			t := s * p[0]
			if t < out {
				out = t
			}
		}
	}
	return out
}
