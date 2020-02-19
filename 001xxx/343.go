func numOfSubarrays(arr []int, k int, threshold int) int {
	threshold *= k
	s := 0
	for i := 0; i < k; i++ {
		s += arr[i]
	}
	o, l := 0, len(arr)
	if s >= threshold {
		o = 1
	}
	for i := k; i < l; i++ {
		s += arr[i] - arr[i-k]
		if s >= threshold {
			o++
		}
	}
	return o
}
