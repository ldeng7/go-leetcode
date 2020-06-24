type Cashier struct {
	m       map[int]int
	c, n, d int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	ca := Cashier{nil, 0, n, 100 - discount}
	m := map[int]int{}
	for i, prod := range products {
		m[prod] = prices[i]
	}
	ca.m = m
	return ca
}

func (this *Cashier) GetBill(products []int, amounts []int) float64 {
	s := 0
	for i, prod := range products {
		s += this.m[prod] * amounts[i]
	}
	this.c++
	if this.c%this.n == 0 {
		return float64(s*this.d) / 100.0
	}
	return float64(s)
}
