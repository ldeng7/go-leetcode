import "sort"

func fourSum(nums []int, target int) [][]int {
	o, l := [][]int{}, len(nums)
	sort.Ints(nums)
	for i := 0; i < l-3; i++ {
		ni := nums[i]
		if i != 0 && ni == nums[i-1] {
			continue
		}
		if ni+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if ni+nums[l-1]+nums[l-2]+nums[l-2] < target {
			continue
		}
		for j := i + 1; j < l-2; j++ {
			nj := nums[j]
			if j != i+1 && nj == nums[j-1] {
				continue
			}
			if ni+nj+nums[j+1]+nums[j+2] > target {
				break
			}
			if ni+nj+nums[l-1]+nums[l-2] < target {
				continue
			}
			b, e := j+1, l-1
			for b < e {
				t := ni + nj + nums[b] + nums[e]
				if t == target {
					o, b = append(o, []int{ni, nj, nums[b], nums[e]}), b+1
					for b < e && nums[b] == nums[b-1] {
						b++
					}
				} else if t > target {
					e--
				} else {
					b++
				}
			}
		}
	}
	return o
}
