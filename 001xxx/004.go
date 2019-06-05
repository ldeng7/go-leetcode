func longestOnes(A []int, K int) int {
	o, i, j := 0, 0, 0
	for i < len(A) && j < len(A) {
		if A[j] == 1 {
			j++
		} else if K != 0 {
			j++
			K--
		} else if A[i] == 1 {
			i++
		} else {
			i++
			j++
		}
		if j-i > o {
			o = j - i
		}
	}
	return o
}
