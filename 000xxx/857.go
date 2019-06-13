import (
	"math"
	"sort"
)

type PriorityQueue struct {
	Arr    []float64
	lessCb func(a, b float64) bool
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.lessCb(pq.Arr[j], pq.Arr[i]) {
			break
		}
		pq.Arr[i], pq.Arr[j] = pq.Arr[j], pq.Arr[i]
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
		if j2 := j1 + 1; j2 < n && pq.lessCb(pq.Arr[j2], pq.Arr[j1]) {
			j = j2
		}
		if !pq.lessCb(pq.Arr[j], pq.Arr[i]) {
			break
		}
		pq.Arr[i], pq.Arr[j] = pq.Arr[j], pq.Arr[i]
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue) Init(arr []float64, lessCb func(float64, float64) bool) *PriorityQueue {
	pq.Arr = arr
	pq.lessCb = lessCb
	if nil == lessCb {
		pq.lessCb = func(a, b float64) bool { return a < b }
	}
	l := len(pq.Arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

func (pq *PriorityQueue) Push(item float64) {
	pq.Arr = append(pq.Arr, item)
	pq.up(len(pq.Arr) - 1)
}

func (pq *PriorityQueue) Pop() (float64, bool) {
	i := len(pq.Arr) - 1
	if i >= 0 {
		pq.Arr[0], pq.Arr[i] = pq.Arr[i], pq.Arr[0]
		pq.down(0, i)
		v := pq.Arr[i]
		pq.Arr = pq.Arr[:i]
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
	q := (&PriorityQueue{}).Init(nil, nil)
	for _, p := range ps {
		s += p[1]
		q.Push(-p[1])
		if len(q.Arr) > K {
			qu, _ := q.Pop()
			s += qu
		}
		if len(q.Arr) == K {
			t := s * p[0]
			if t < out {
				out = t
			}
		}
	}
	return out
}
