import "math"

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	m := make([]bool, n)
	m[0], m[1] = true, true
	e := int(math.Sqrt(float64(n)))
	for i := 2; i <= e; i++ {
		if !m[i] {
			for j := i * i; j < n; j += i {
				m[j] = true
			}
		}
	}
	o := 0
	for _, b := range m {
		if !b {
			o++
		}
	}
	return o
}
