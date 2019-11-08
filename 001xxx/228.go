func missingNumber(arr []int) int {
	l := len(arr)
	o := ((arr[0] + arr[l-1]) * (l + 1)) >> 1
	for _, a := range arr {
		o -= a
	}
	return o
}
