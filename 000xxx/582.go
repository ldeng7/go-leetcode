func killProcess(pid []int, ppid []int, kill int) []int {
	out := []int{}
	m := map[int][]int{}
	for i := 0; i < len(pid); i++ {
		m[ppid[i]] = append(m[ppid[i]], pid[i])
	}
	var cal func(int)
	cal = func(i int) {
		out = append(out, i)
		for _, j := range m[i] {
			cal(j)
		}
	}
	cal(kill)
	return out
}
