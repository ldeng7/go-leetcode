func minJumps(arr []int) int {
	o, l := 0, len(arr)
	m := map[int][]int{}
	q := make([]int, 1, 32)
	v := make([]bool, l)
	for i, a := range arr {
		m[a] = append(m[a], i)
	}
	v[0] = true

	for 0 != len(q) {
		for i := len(q); i > 0; q, i = q[1:], i-1 {
			f := q[0]
			if f == l-1 {
				return o
			}
			for _, j := range m[arr[f]] {
				if !v[j] {
					q, v[j], arr[j] = append(q, j), true, -100000001
				}
			}
			if f < l-1 && !v[f+1] {
				q, v[f+1] = append(q, f+1), true
			}
			if f > 0 && !v[f-1] {
				q, v[f-1] = append(q, f-1), true
			}
		}
		o++
	}
	return o
}
