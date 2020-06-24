import "fmt"

func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func simplifiedFractions(n int) []string {
	o := make([]string, 0, 64)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				o = append(o, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return o
}
