func numFriendRequests(ages []int) int {
	ns, ss := make([]int, 121), make([]int, 121)
	for _, age := range ages {
		ns[age]++
	}
	for i := 1; i < 121; i++ {
		ss[i] = ns[i] + ss[i-1]
	}
	out := 0
	for i := 15; i < 121; i++ {
		if ns[i] == 0 {
			continue
		}
		out += (ss[i]-ss[i>>1+7])*ns[i] - ns[i]
	}
	return out
}
