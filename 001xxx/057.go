import "sort"

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

type WorkerBike struct {
	WorkerId int
	BikeId   int
	Dist     int
}

func less(a, b *WorkerBike) bool {
	if a.Dist != b.Dist {
		return a.Dist < b.Dist
	} else if a.WorkerId != b.WorkerId {
		return a.WorkerId < b.WorkerId
	}
	return a.BikeId < b.BikeId
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	lw, lb := len(workers), len(bikes)
	wbs := make([]WorkerBike, lw*lb)
	out := make([]int, lw)
	for i := 0; i < lw; i++ {
		out[i] = -1
	}
	b := make([]bool, lb)

	k := 0
	for i := 0; i < lw; i++ {
		w := workers[i]
		for j := 0; j < lb; j++ {
			b := bikes[j]
			wbs[k] = WorkerBike{i, j, abs(w[0]-b[0]) + abs(w[1]-b[1])}
			k++
		}
	}
	sort.Slice(wbs, func(i, j int) bool {
		return less(&wbs[i], &wbs[j])
	})

	for _, wb := range wbs {
		if out[wb.WorkerId] == -1 && !b[wb.BikeId] {
			out[wb.WorkerId], b[wb.BikeId] = wb.BikeId, true
		}
	}
	return out
}
