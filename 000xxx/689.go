import "math"

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	l := len(nums)
	sum, sums := 0, make([]int, l+1)
	for i, n := range nums {
		sum += n
		sums[i+1] = sum
	}
	ls, rs := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		rs[i] = l - k
	}

	for i, t := k, sums[k]-sums[0]; i < l; i++ {
		if sums[i+1]-sums[i+1-k] > t {
			ls[i] = i + 1 - k
			t = sums[i+1] - sums[i+1-k]
		} else {
			ls[i] = ls[i-1]
		}
	}
	for i, t := l-1-k, sums[l]-sums[l-k]; i >= 0; i-- {
		if sums[i+k]-sums[i] >= t {
			rs[i] = i
			t = sums[i+k] - sums[i]
		} else {
			rs[i] = rs[i+1]
		}
	}
	m := math.MinInt64
	var o []int
	for i := k; i <= l-k<<1; i++ {
		b, e := ls[i-1], rs[i+k]
		t := (sums[i+k] - sums[i]) + (sums[b+k] - sums[b]) + (sums[e+k] - sums[e])
		if m < t {
			m, o = t, []int{b, i, e}
		}
	}
	return o
}
