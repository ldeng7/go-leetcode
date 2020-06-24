import "math"

var m = map[int]int{}

func cal(n int) int {
	c := 0
	for n&1 == 0 {
		c, n = c+1, n>>1
	}
	if c > 3 || c == 2 {
		return 0
	} else if c == 3 {
		if n > 1 {
			return 0
		}
		return 15
	}
	a, b, p := c, 3, 2
	for n > 2 {
		pn := n
		for i := b; i <= int(math.Sqrt(float64(n))); i += 2 {
			if n%i != 0 {
				continue
			} else if a == 1 {
				return 0
			}
			for n%i == 0 {
				c, n = c+1, n/i
			}
			if c > 3 || c == 2 {
				return 0
			} else if c == 3 {
				if n > 1 {
					return 0
				}
				return i*i*i + i*i + i + 1
			}
			a, b, p = c, i+2, i
			break
		}
		if n == pn {
			break
		}
	}
	if a == 1 && n > p {
		return p*n + p + n + 1
	}
	return 0
}

func sumFourDivisors(nums []int) int {
	o := 0
	for _, n := range nums {
		v, ok := m[n]
		if !ok {
			v = cal(n)
			m[n] = v
		}
		o += v
	}
	return o
}
