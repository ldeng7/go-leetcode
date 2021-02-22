func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func specialArray(nums []int) int {
	l := len(nums)
	ar := make([]int, l+1)
	for _, num := range nums {
		ar[min(num, l)]++
	}
	for i := l; i >= 0; i-- {
		if i < l {
			ar[i] += ar[i+1]
		}
		if ar[i] == i {
			return i
		}
	}
	return -1
}
