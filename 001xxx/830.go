const MOD = 1000000007

var ar [3001]int

func init() {
	ar[1] = 1
	for i := 2; i <= 3000; i++ {
		ar[i] = ar[MOD%i] * (MOD - MOD/i) % MOD
	}
}

func makeStringSorted(s string) int {
	o, l, t := 0, len(s), 1
	cs := [26]int{}
	for i := l - 1; i >= 0; i-- {
		c := int(s[i] - 'a')
		cs[c]++
		t = t * ar[cs[c]] % MOD
		a := 0
		for j := 0; j < c; j++ {
			a += cs[j]
		}
		o = (o + t*a) % MOD
		t = t * (l - i) % MOD
	}
	return o
}
