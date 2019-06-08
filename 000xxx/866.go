import (
	"math"
	"strconv"
)

func primePalindrome(N int) int {
	if N >= 8 && N <= 11 {
		return 11
	}
	for i := 1; i < 1e5; i++ {
		s := strconv.Itoa(i)
		t := []byte(s)
		l := len(t)
		for j := 0; j < l>>1; j++ {
			t[j], t[l-j-1] = t[l-j-1], t[j]
		}
		k, _ := strconv.Atoi(s + string(t[1:]))
		if k >= N && isPrime(k) {
			return k
		}
	}
	return -1
}

func isPrime(n int) bool {
	if n < 2 || n&1 == 0 {
		return n == 2
	}
	s := int(math.Sqrt(float64(n)))
	for i := 3; i <= s; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
