var out = [][]int{}

func do(candidates []int, target int, reg []int) {
	sum := 0
	for _, num := range reg {
		sum += num
		if sum > target {
			return
		}
	}
	if sum == target {
		reg1 := make([]int, len(reg))
		copy(reg1, reg)
		out = append(out, reg1)
		return
	}
	for i, num := range candidates {
		reg1 := append(reg, num)
		do(candidates[i:], target, reg1)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	out = [][]int{}
	do(candidates, target, []int{})
	return out
}