import "sort"

func permuteUnique(nums []int) [][]int {
	o := [][]int{}
	sort.Ints(nums)
	var permute func(int)
	permute = func(b int) {
		if b >= len(nums) {
			nums1 := make([]int, len(nums))
			copy(nums1, nums)
			o = append(o, nums1)
		}
		for i := b; i < len(nums); i++ {
			j := i - 1
			for j >= b && nums[j] != nums[i] {
				j--
			}
			if j != b-1 {
				continue
			}
			nums[i], nums[b] = nums[b], nums[i]
			permute(b + 1)
			nums[i], nums[b] = nums[b], nums[i]
		}
	}
	permute(0)
	return o
}
