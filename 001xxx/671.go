import "sort"

func minimumMountainRemovals(nums []int) int {
	l := len(nums)
	ar1, ar2 := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		ar1[i], ar2[i] = 1, 1
	}

	t := make([]int, 0, 16)
	for i, n := range nums {
		j := sort.Search(len(t), func(k int) bool {
			return t[k] >= n
		})
		if j == len(t) {
			t = append(t, n)
		} else {
			t[j] = n
		}
		ar1[i] = len(t)
	}

	t = make([]int, 0, 16)
	for i := l - 1; i >= 0; i-- {
		n := nums[i]
		j := sort.Search(len(t), func(k int) bool {
			return t[k] >= n
		})
		if j == len(t) {
			t = append(t, n)
		} else {
			t[j] = n
		}
		ar2[i] = len(t)
	}

	o := 0
	for i := 1; i < l-1; i++ {
		if c1, c2 := ar1[i], ar2[i]; c1 > 1 && c2 > 1 {
			if t := c1 + c2 - 1; t > o {
				o = t
			}
		}
	}
	return l - o
}
