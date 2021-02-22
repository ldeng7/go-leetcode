func getWinner(arr []int, k int) int {
	i := 0
	for t := 0; t < k && i < len(arr)-1; i++ {
		if a := arr[i]; a > arr[i+1] {
			arr[i+1] = a
			t++
		} else {
			t = 1
		}
	}
	return arr[i]
}
