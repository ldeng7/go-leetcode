type StockSpanner struct {
	st [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{}}
}

func (this *StockSpanner) Next(price int) int {
	o := 1
	for 0 != len(this.st) && this.st[len(this.st)-1][0] <= price {
		e := len(this.st) - 1
		o += this.st[e][1]
		this.st = this.st[:e]
	}
	this.st = append(this.st, [2]int{price, o})
	return o
}
