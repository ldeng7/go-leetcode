func isPossible(nums []int) bool {
	cnts, reqs := map[int]int{}, map[int]int{}
	for _, n := range nums {
		cnts[n]++
	}
	for _, n := range nums {
		if cnts[n] == 0 {
			continue
		} else if reqs[n] > 0 {
			reqs[n]--
			reqs[n+1]++
		} else if cnts[n+1] > 0 && cnts[n+2] > 0 {
			cnts[n+1]--
			cnts[n+2]--
			reqs[n+3]++
		} else {
			return false
		}
		cnts[n]--
	}
	return true
}
