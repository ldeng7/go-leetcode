func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	l, r := n-1, n-1
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			l = i - 1
			break
		}
	}
	if l == n-1 {
		return 0
	}
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			r = i + 1
			break
		}
	}
	a, b := 0, r
	o := min(n-l-1, r)
	for a <= l {
		if b == n || arr[a] <= arr[b] {
			o = min(o, b-a-1)
			a++
		} else {
			b++
		}
	}
	return o
}
