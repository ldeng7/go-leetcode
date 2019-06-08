import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	out := [][]int{}
	n := min(k, len(nums1)*len(nums2))
	t := make([]int, len(nums1))
	for i := 0; i < n; i++ {
		cur, sum := 0, math.MaxInt64
		for j := 0; j < len(nums1); j++ {
			if t[j] < len(nums2) {
				s1 := nums1[j] + nums2[t[j]]
				if sum >= s1 {
					cur, sum = j, s1
				}
			}
		}
		out = append(out, []int{nums1[cur], nums2[t[cur]]})
		t[cur]++
	}
	return out
}
