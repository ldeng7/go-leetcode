func combinationSum3(k int, n int) [][]int {
	out, ar := [][]int{}, []int{}
	var cal func(n, level int)
	cal = func(n, level int) {
		if n < 0 {
			return
		}
		if n == 0 && len(ar) == k {
			ar1 := make([]int, len(ar))
			copy(ar1, ar)
			out = append(out, ar1)
		}
		for i := level; i <= 9; i++ {
			ar = append(ar, i)
			cal(n-i, i+1)
			ar = ar[:len(ar)-1]
		}
	}
	cal(n, 1)
	return out
}
