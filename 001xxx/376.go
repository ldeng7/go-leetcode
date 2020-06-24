func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	ts := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ts[i] = -1
	}
	ts[headID] = 0
	o := 0

	var cal func(int)
	cal = func(i int) {
		if -1 != ts[i] {
			return
		}
		m := manager[i]
		cal(m)
		ts[i] = ts[m] + informTime[m]
	}

	for i := 0; i < n; i++ {
		if -1 == manager[i] {
			ts[i] = 0
		} else {
			cal(i)
			if t := ts[i]; t > o {
				o = t
			}
		}
	}
	return o
}
