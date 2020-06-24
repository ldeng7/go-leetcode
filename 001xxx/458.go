func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	o, m, n := -100000, len(nums1), len(nums2)
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		ar := make([]int, n+1)
		for j := 0; j <= n; j++ {
			ar[j] = -100000
		}
		t[i] = ar
	}
	for i := m - 1; i >= 0; i-- {
		ar, ar1, num1 := t[i], t[i+1], nums1[i]
		for j := n - 1; j >= 0; j-- {
			ar[j] = max(max(ar[j+1], ar1[j]), num1*nums2[j]+max(ar1[j+1], 0))
			o = max(o, ar[j])
		}
	}
	return o
}
