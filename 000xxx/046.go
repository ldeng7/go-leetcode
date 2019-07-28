func permute(nums []int) [][]int {
	if 1 == len(nums) {
		return [][]int{{nums[0]}}
	}
	o := [][]int{}
	for i, num := range nums {
		r, k := make([]int, len(nums)-1), 0
		for j, num1 := range nums {
			if i != j {
				r[k], k = num1, k+1
			}
		}
		for _, ar := range permute(r) {
			ar1 := make([]int, len(ar)+1)
			ar1[0] = num
			copy(ar1[1:], ar)
			o = append(o, ar1)
		}
	}
	return o
}
