import "sort"

func permuteUnique(nums []int) [][]int {
	out := [][]int{}
	sort.Ints(nums)
	var permute func(int)
	permute = func(start int) {
		if start >= len(nums) {
			nums1 := make([]int, len(nums))
			copy(nums1, nums)
			out = append(out, nums1)
		}
		for i := start; i < len(nums); i++ {
			j := i - 1
			for j >= start && nums[j] != nums[i] {
				j--
			}
			if j != start-1 {
				continue
			}
			nums[i], nums[start] = nums[start], nums[i]
			permute(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	permute(0)
	return out
}
