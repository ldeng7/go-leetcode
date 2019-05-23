import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	i0 := sort.Search(len(A), func(j int) bool {
		return A[j] >= 0
	})
	i, sum := 0, 0
	for ; i < K && i < i0; i++ {
		sum -= A[i]
	}
	if K > i0 && (K-i0)&1 == 1 {
		if i > 0 && A[i] > -A[i-1] {
			sum += A[i-1] * 2
		} else {
			sum -= A[i]
			i++
		}
	}
	for ; i < len(A); i++ {
		sum += A[i]
	}
	return sum
}
