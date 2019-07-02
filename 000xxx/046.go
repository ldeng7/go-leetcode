func permute(nums []int) [][]int {
	if 1 == len(nums) {
		return [][]int{{nums[0]}}
	}
	out := [][]int{}
	for i, num := range nums {
		rest, k := make([]int, len(nums)-1), 0
		for j, num1 := range nums {
			if i != j {
				rest[k], k = num1, k+1
			}
		}
		for _, ar := range permute(rest) {
			ar1 := make([]int, len(ar)+1)
			ar1[0] = num
			copy(ar1[1:], ar)
			out = append(out, ar1)
		}
	}
	return out
}
