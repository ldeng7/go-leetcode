import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const N = 1001000

var t [N]int
var f [N]int
var o [N / 10]int

func init() {
	t[0], t[1] = 1, 1
	primes := [N]int{}
	c := 0
	for i := 2; i < N; i++ {
		if 0 == t[i] {
			primes[c], t[i] = i, i
			c++
		}
		j := 0
		for {
			p := primes[j]
			k := p * i
			if k >= N {
				break
			}
			t[k] = p
			if 0 == i%p {
				break
			}
			j++
		}
	}
}

func splitArray(nums []int) int {
	l := len(nums) - 1
	for i := 0; i < N; i++ {
		f[i] = math.MaxInt32
	}
	x := nums[0]
	for 1 != x {
		tx := t[x]
		f[tx], x = 1, x/tx
	}

	for i, x := range nums {
		m := math.MaxInt32
		for 1 != x {
			tx := t[x]
			m, x = min(m, f[tx]), x/tx
		}
		o[i] = m
		if i < l {
			x = nums[i+1]
			for 1 != x {
				tx := t[x]
				f[tx], x = min(m+1, f[tx]), x/tx
			}
		}
	}
	return o[l]
}
