func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		out := 0
		for i := 0; i < n; i++ {
			if knows(out, i) {
				out = i
			}
		}
		for i := 0; i < out; i++ {
			if knows(out, i) || (!knows(i, out)) {
				return -1
			}
		}
		for i := out + 1; i < n; i++ {
			if !knows(i, out) {
				return -1
			}
		}
		return out
	}
}
