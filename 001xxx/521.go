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
	o, t := 10000000, arr[0]
	for _, a := range arr {
		if t < target {
			t = a
		}
		t &= a
		o = min(o, abs(target-t))
	}
	return o
}
