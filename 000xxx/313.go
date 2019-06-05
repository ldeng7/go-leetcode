import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = 1
	}
	indices := make([]int, len(primes))
	for i := 1; i < n; i++ {
		t[i] = math.MaxInt64
		for j, p := range primes {
			if t1 := p * t[indices[j]]; t1 < t[i] {
				t[i] = t1
			}
		}
		for j, p := range primes {
			if t[i] == p*t[indices[j]] {
				indices[j]++
			}
		}
	}
	return t[n-1]
}
