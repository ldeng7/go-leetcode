func sumOddLengthSubarrays(arr []int) int {
	o, l := 0, len(arr)
	for i, a := range arr {
		l, r := i+1, l-i
		o += (((l+1)>>1)*((r+1)>>1) + (l>>1)*(r>>1)) * a
	}
	return o
}
