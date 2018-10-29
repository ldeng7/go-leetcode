import "sort"

func lengthOfLIS(nums []int) int {
	t := []int{}
	for _, n := range nums {
		j := sort.Search(len(t), func(i int) bool {
			return t[i] >= n
		})
		if j == len(t) {
			t = append(t, n)
		} else {
			t[j] = n
		}
	}
	return len(t)
}
