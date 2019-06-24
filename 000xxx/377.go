import "sort"

func combinationSum4(nums []int, target int) int {
	t := make([]int, target+1)
	t[0] = 1
	sort.Ints(nums)
	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if i < n {
				break
			}
			t[i] += t[i-n]
		}
	}
	return t[target]
}
