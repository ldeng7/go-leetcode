import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})
	o, c := 0, 0
	for _, t := range tasks {
		a, m := t[0], t[1]
		o += max(0, max(a, m)-c)
		c = max(a, max(c, m)) - a
	}
	return o
}
