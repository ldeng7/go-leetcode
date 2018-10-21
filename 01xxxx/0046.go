func permute(nums []int) [][]int {
	if 1 == len(nums) {
		return [][]int{{nums[0]}}
	}
	out := [][]int{}
	for i, num := range nums {
		rest := []int{}
		for j, num1 := range nums {
			if i != j {
				rest = append(rest, num1)
			}
		}
		for _, ar := range permute(rest) {
			ar1 := []int{num}
			ar1 = append(ar1, ar...)
			out = append(out, ar1)
		}
	}
	return out
}
