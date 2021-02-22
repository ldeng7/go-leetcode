func sumOfUnique(nums []int) int {
	t := [101]int{}
	for _, n := range nums {
		t[n]++
	}
	o := 0
	for i := 0; i <= 100; i++ {
		if t[i] == 1 {
			o += i
		}
	}
	return o
}
