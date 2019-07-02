func duplicateZeros(arr []int) {
	l := len(arr)
	t := make([]int, l<<1)
	for i, j := 0, 0; i < l; i++ {
		a := arr[i]
		if a == 0 {
			t[j], t[j+1], j = 0, 0, j+2
		} else {
			t[j], j = a, j+1
		}
	}
	for i := 0; i < l; i++ {
		arr[i] = t[i]
	}
}
