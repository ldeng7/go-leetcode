func subsets(nums []int) [][]int {
	out := [][]int{{}}
	for _, num := range nums {
		l := len(out)
		for i := 0; i < l; i++ {
			ar := out[i]
			arNew := make([]int, len(ar)+1)
			copy(arNew, ar)
			arNew[len(ar)] = num
			out = append(out, arNew)
		}
	}
	return out
}
