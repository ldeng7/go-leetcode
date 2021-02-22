func checkPartitioning(s string) bool {
	cal := func(i, j int) bool {
		for ; i < j && s[i] == s[j]; i, j = i+1, j-1 {
		}
		return i >= j
	}

	l := len(s)
	arl, arr := make([]int, 1, 16), make([]int, 1, 16)
	arl[0], arr[0] = 0, l-1
	for i := 1; i < l-2; i++ {
		if cal(0, i) {
			arl = append(arl, i)
		}
	}
	for i := l - 2; i > 1; i-- {
		if cal(i, l-1) {
			arr = append(arr, i)
		}
	}

	n, j := len(arr), 0
	for _, i := range arl {
		for ; j < n && i+2 > arr[n-j-1]; j++ {
		}
		if j >= n {
			break
		} else if cal(i+1, arr[n-j-1]-1) {
			return true
		}
	}
	return false
}
