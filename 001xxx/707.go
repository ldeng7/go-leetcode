import "sort"

func maximizeXor(nums []int, queries [][]int) []int {
	sort.Ints(nums)
	o := make([]int, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		if b < nums[0] {
			o[i] = -1
			continue
		}

		l, r := 0, len(nums)-1
		for l < r {
			m := (l + r + 1) >> 1
			if nums[m] <= b {
				l = m
			} else {
				r = m - 1
			}
		}

		l, v := 0, 0
		for j := 30; j >= 0; j-- {
			if (a>>j)&1 == 1 {
				if (nums[l]>>j)&1 == 1 {
					v ^= (1 << j)
					continue
				}
				l1, r1 := l, r
				for l1 < r1 {
					m := (l1 + r1 + 1) >> 1
					if (nums[m]>>j)&1 == 0 {
						l1 = m
					} else {
						r1 = m - 1
					}
				}
				r = r1
			} else {
				if (nums[r]>>j)&1 == 0 {
					continue
				}
				l1, r1 := l, r
				v ^= (1 << j)
				for l1 < r1 {
					m := (l1 + r1) >> 1
					if (nums[m]>>j)&1 == 0 {
						l1 = m + 1
					} else {
						r1 = m
					}
				}
				l = l1
			}
		}
		o[i] = v ^ a
	}
	return o
}
