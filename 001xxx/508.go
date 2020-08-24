var buckets = [100001]int{}

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, n)
	sums[0] = nums[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	for i := 0; i <= sums[n-1]; i++ {
		buckets[i] = 0
	}
	for i, s := range sums {
		buckets[s]++
		for j := 0; j < i; j++ {
			buckets[s-sums[j]]++
		}
	}
	ar := make([]int, 0, n*(n+1)/2)
	for i, b := range buckets {
		for b > 0 {
			ar = append(ar, i)
			b--
			buckets[i]--
		}
	}
	o := 0
	for i := left - 1; i < right; i++ {
		o = (o + ar[i]) % 1000000007
	}
	return o
}
