func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	m := map[int]int{}
	var cal func(int) int
	cal = func(tar int) int {
		if v, ok := m[tar]; ok {
			return v
		} else if tar == 0 {
			return 0
		} else if tar == 1 {
			return inc
		}
		o := inc * tar
		for i, v := range jump {
			t, k := tar/v, tar%v
			if k == 0 {
				o = min(o, cal(t)+cost[i])
			} else {
				o = min(o, cal(t)+cost[i]+k*inc)
				o = min(o, cal(t+1)+cost[i]+(v-k)*dec)
			}
		}
		m[tar] = o
		return o
	}
	return cal(target) % 1000000007
}
