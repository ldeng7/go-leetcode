func waysToSplit(nums []int) int {
	o, l := 0, len(nums)
	for i := 1; i < l; i++ {
		nums[i] += nums[i-1]
	}
	for i, j, k := 0, 0, 0; i < l-2; i++ {
		for ; j <= i || (j < l-1 && nums[j] < nums[i]*2); j++ {
		}
		for ; k < j || (k < l-1 && nums[k]-nums[i] <= nums[l-1]-nums[k]); k++ {
		}
		o = (o + k - j) % 1000000007
	}
	return o
}
