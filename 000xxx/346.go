type MovingAverage struct {
	ar   []int
	c, i int
	sum  int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{[]int{}, size, 0, 0}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.ar) == this.c {
		this.ar[this.i], this.sum = val, this.sum-this.ar[this.i]+val
		this.i++
		if this.i == this.c {
			this.i = 0
		}
		return float64(this.sum) / float64(this.c)
	} else {
		this.ar = append(this.ar, val)
		this.sum += val
		return float64(this.sum) / float64(len(this.ar))
	}
}
