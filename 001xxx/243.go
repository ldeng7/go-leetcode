func transformArray(arr []int) []int {
	l := len(arr)
	if l <= 2 {
		return arr
	}
	o := make([]int, l)
	copy(o, arr)
	con := true
	for con {
		con = false
		a, b, c := 0, arr[0], arr[1]
		for i := 1; i < l-1; i++ {
			a, b, c = b, c, arr[i+1]
			if b < a && b < c {
				o[i]++
				con = true
			}
			if b > a && b > c {
				o[i]--
				con = true
			}
		}
		copy(arr, o)
	}
	return o
}
