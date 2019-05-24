import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.IntSlice(nums).Sort()
	out := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, len(nums)-1
			for k < l {
				s := nums[i] + nums[j] + nums[k] + nums[l]
				if s == target {
					ok := true
					for _, o := range out {
						if o[0] == nums[i] && o[1] == nums[j] && o[2] == nums[k] && o[3] == nums[l] {
							ok = false
							break
						}
					}
					if ok {
						out = append(out, []int{nums[i], nums[j], nums[k], nums[l]})
					}
					k++
					l--
				} else if s < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return out
}
