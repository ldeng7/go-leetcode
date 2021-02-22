const M = 1000000007

func power(a, b int) int {
	x, y := 1, a
	for ; b > 0; b >>= 1 {
		if b&1 == 1 {
			x = x * y % M
		}
		y = y * y % M
	}
	return x
}

type Fancy struct {
	a, m int
	ar   [][3]int
}

func Constructor() Fancy {
	return Fancy{0, 1, make([][3]int, 0, 16)}
}

func (this *Fancy) Append(val int) {
	this.ar = append(this.ar, [3]int{val, this.a, this.m})
}

func (this *Fancy) AddAll(inc int) {
	this.a += inc
}

func (this *Fancy) MultAll(m int) {
	this.m = this.m * m % M
	this.a = this.a * m % M
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.ar) {
		return -1
	}
	t := this.ar[idx]
	m := this.m * power(t[2], M-2) % M
	a := (this.a - (t[1]*m)%M + M) % M
	return (m*t[0] + a) % M
}
