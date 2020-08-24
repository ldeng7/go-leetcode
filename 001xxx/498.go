import "sort"

const MOD = 1000000007

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ar := make([]int, n)
	ar[0] = 1
	for i := 1; i < n; i++ {
		ar[i] = (ar[i-1] << 1) % MOD
	}

	o, l, r := 0, 0, n-1
	for l <= r {
		if nums[l]+nums[r] <= target {
			o = (o + ar[r-l]) % MOD
			l++
		} else {
			r--
		}
	}
	return o
}
