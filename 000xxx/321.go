func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxVector(nums []int, k int) []int {
	o, n := []int{}, len(nums)-k
	for _, num := range nums {
		for n > 0 && 0 != len(o) && o[len(o)-1] < num {
			o, n = o[:len(o)-1], n-1
		}
		o = append(o, num)
	}
	return o[:k]
}

func cmpArr(nums1, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	l := min(l1, l2)
	for i := 0; i < l; i++ {
		a, b := nums1[i], nums2[i]
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	if l1 > l2 {
		return 1
	} else if l1 < l2 {
		return -1
	}
	return 0
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	o, l1, l2 := []int{}, len(nums1), len(nums2)
	ie := min(k, l1)
	for i := max(0, k-l2); i <= ie; i++ {
		m, m1, m2 := []int{}, maxVector(nums1, i), maxVector(nums2, k-i)
		for 0 != len(m1) || 0 != len(m2) {
			if cmpArr(m1, m2) == 1 {
				m = append(m, m1[0])
				m1 = m1[1:]
			} else {
				m = append(m, m2[0])
				m2 = m2[1:]
			}
		}
		if cmpArr(m, o) == 1 {
			o = m
		}
	}
	return o
}
