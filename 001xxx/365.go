func smallerNumbersThanCurrent(nums []int) []int {
	cnts := [101]int{}
	for _, n := range nums {
		cnts[n]++
	}
	for i := 1; i < 100; i++ {
		cnts[i] += cnts[i-1]
	}
	o := make([]int, len(nums))
	for i, n := range nums {
		if 0 != n {
			o[i] = cnts[n-1]
		} else {
			o[i] = 0
		}
	}
	return o
}
