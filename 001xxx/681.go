import "sort"

func minimumIncompatibility(nums []int, k int) int {
	var o, n int = 1e10, len(nums)
	var cal func(int, int, int, int, int, int, int)
	cal = func(i, c, l, h, v, s, p int) {
		if t := s + h - l; t >= o {
			return
		} else if c == n/k {
			j := 0
			for ; j < n && (v>>j)&1 == 1; j++ {
			}
			if j == n {
				o = t
			} else {
				nj := nums[j]
				cal(j, 1, nj, nj, v|(1<<j), t, nj)
			}
		} else {
			for j := i + 1; j < n; j++ {
				if nj := nums[j]; nj != p && (v>>j)&1 == 0 {
					p = nj
					cal(j, c+1, l, nj, v|(1<<j), s, nj)
				}
			}
		}
	}

	fs := [17]int{}
	for _, num := range nums {
		fs[num]++
		if fs[num] > k {
			return -1
		}
	}
	sort.Ints(nums)
	n0 := nums[0]
	cal(0, 1, n0, n0, 1, 0, n0)
	return o
}
