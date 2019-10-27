import "sort"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

type Job struct {
	s int
	e int
	p int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	l := len(profit)
	jobs := make([]Job, l)
	for i := 0; i < l; i++ {
		jobs[i] = Job{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].e < jobs[j].e
	})

	t := make([]int, l)
	t[0] = jobs[0].p
	for i := 1; i < l; i++ {
		k, s := 0, jobs[i].s
		for j := i - 1; j >= 0; j-- {
			if jobs[j].e <= s {
				k = t[j]
				break
			}
		}
		t[i] = max(t[i-1], k+jobs[i].p)
	}
	return t[l-1]
}
