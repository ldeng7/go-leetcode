func maxJumps(arr []int, d int) int {
	o, l := 0, len(arr)
	t := make([]int, l)
	var cal func(int) int
	cal = func(i int) int {
		if v := t[i]; 0 != v {
			return v
		}
		o, a := 1, arr[i]
		for j := i - 1; j >= 0 && i-j <= d && arr[j] < a; j-- {
			if v := 1 + cal(j); v > o {
				o = v
			}
		}
		for j := i + 1; j < l && j-i <= d && arr[j] < a; j++ {
			if v := 1 + cal(j); v > o {
				o = v
			}
		}
		t[i] = o
		return o
	}

	for i := 0; i < l; i++ {
		if v := cal(i); v > o {
			o = v
		}
	}
	return o
}
