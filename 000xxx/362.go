type HitCounter struct {
	ts [300]int
	cs [300]int
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (this *HitCounter) Hit(timestamp int) {
	i := timestamp % 300
	if this.ts[i] != timestamp {
		this.ts[i] = timestamp
		this.cs[i] = 1
	} else {
		this.cs[i]++
	}
}

func (this *HitCounter) GetHits(timestamp int) int {
	tb := timestamp - 299
	sum := 0
	for i, t := range this.ts {
		if t >= tb {
			sum += this.cs[i]
		}
	}
	return sum
}
