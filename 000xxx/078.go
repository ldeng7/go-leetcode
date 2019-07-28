func subsets(nums []int) [][]int {
	out := [][]int{{}}
	for _, num := range nums {
		l := len(out)
		for i := 0; i < l; i++ {
			ar := out[i]
			ar1 := make([]int, len(ar)+1)
			copy(ar1, ar)
			ar1[len(ar)] = num
			out = append(out, ar1)
		}
	}
	return out
}
