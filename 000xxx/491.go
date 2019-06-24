func findSubsequences(nums []int) [][]int {
	o, t := [][]int{}, []int{}
	var cal func(int)
	cal = func(i0 int) {
		if len(t) >= 2 {
			t1 := make([]int, len(t))
			copy(t1, t)
			o = append(o, t1)
		}
		s := map[int]bool{}
		for i := i0; i < len(nums); i++ {
			n := nums[i]
			if (0 != len(t) && t[len(t)-1] > n) || s[n] {
				continue
			}
			t, s[n] = append(t, n), true
			cal(i + 1)
			t = t[:len(t)-1]
		}
	}
	cal(0)
	return o
}
