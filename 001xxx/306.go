func canReach(arr []int, start int) bool {
	n := len(arr)
	v := make([]bool, n)
	q := make([]int, 1, 32)
	q[0] = start
	for 0 != len(q) {
		t := q[0]
		q = q[1:]
		v[t] = true
		num := arr[t]
		if 0 == num {
			return true
		}
		l, r := t-num, t+num
		if l >= 0 && l < n && !v[l] {
			q = append(q, l)
		}
		if r >= 0 && r < n && !v[r] {
			q = append(q, r)
		}
	}
	return false
}
