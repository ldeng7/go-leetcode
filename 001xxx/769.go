func minOperations(boxes string) []int {
	l := len(boxes)
	o := make([]int, l)
	for i, s, c := 0, 0, 0; i < l; i++ {
		s += c
		o[i] = s
		if boxes[i] == '1' {
			c++
		}
	}

	c := 0
	if boxes[l-1] == '1' {
		c = 1
	}
	for i, s := l-2, 0; i >= 0; i-- {
		s += c
		o[i] += s
		if boxes[i] == '1' {
			c++
		}
	}
	return o
}
