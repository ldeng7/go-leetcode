func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func closestToTarget(arr []int, target int) int {
	o := abs(arr[0] - target)
	t := []int{arr[0]}
	for _, a := range arr {
		t1 := make([]int, 1, len(t))
		t1[0] = a
		p := a
		o = min(o, abs(a-target))
		for _, b := range t {
			if c := b & a; c != p {
				t1 = append(t1, c)
				o = min(o, abs(c-target))
				p = c
			}
		}
		t = t1
	}
	return o
}
