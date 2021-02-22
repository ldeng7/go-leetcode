func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	l := len(boxes) - 1
	t := make([]int, l+2)
	t[l] = 2
	c, s := 1, boxes[l][1]
	q := make([]int, 0, 16)
	for i, n := l-1, l; i >= 0; i-- {
		c, s = c+1, s+boxes[i][1]
		if i <= n-1 && boxes[i][0] != boxes[i+1][0] {
			q = append(q, i)
		}
		for c > maxBoxes || s > maxWeight {
			c, s, n = c-1, s-boxes[n][1], n-1
			if 0 != len(q) && q[0] > n-1 {
				q = q[1:]
			}
		}
		t[i] = len(q) + 2 + t[n+1]
		if len(q) >= 1 {
			t[i] = min(t[i], len(q)+1+t[q[0]+1])
		} else {
			t[i] = len(q) + 2 + t[n+1]
		}
	}
	return t[0]
}
