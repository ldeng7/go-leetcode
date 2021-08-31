import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minAbsDifference(nums []int, goal int) int {
	s := 0
	for i, n := range nums {
		if n > 0 {
			s += n
		} else {
			nums[i] = -n
		}
	}
	sort.Ints(nums)

	i, l := 0, len(nums)
	for ; i < l && nums[i] == 0; i++ {
	}
	nums, l = nums[i:], l-i
	d := s - goal
	if d <= 0 {
		return -d
	}

	o := d
	ss := make([]int, l+1)
	for i := l - 1; i >= 0; i-- {
		ss[i] = ss[i+1] + nums[l-i-1]
	}

	var cal func(int, int)
	cal = func(i, j int) {
		o = min(o, abs(j-d))
		if j >= d || o == 0 || i == l {
			return
		} else if t := d - (j + ss[i]); t >= 0 {
			o = min(o, t)
			return
		}
		cal(i+1, j+nums[l-i-1])
		cal(i+1, j)
	}
	cal(0, 0)
	return o
}
