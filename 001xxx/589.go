import "sort"

func maxSumRangeQuery(nums []int, requests [][]int) int {
	l := len(nums)
	ar := make([]int, l+1)
	for _, r := range requests {
		a, b := r[0], r[1]
		ar[a]++
		ar[b+1]--
	}
	for i := 1; i < l; i++ {
		ar[i] += ar[i-1]
	}
	sort.Ints(ar[:l])
	sort.Ints(nums)
	o := 0
	for i := l - 1; i >= 0 && nums[i] > 0; i-- {
		o = (o + ar[i]*nums[i]) % 1000000007
	}
	return o
}
