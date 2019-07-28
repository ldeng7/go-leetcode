type TwoSum struct {
	m map[int]int
}

func Constructor() TwoSum {
	return TwoSum{map[int]int{}}
}

func (this *TwoSum) Add(number int) {
	this.m[number]++
}

func (this *TwoSum) Find(value int) bool {
	for n, c := range this.m {
		n1 := value - n
		if n != n1 {
			if c1 := this.m[n1]; c1 > 0 {
				return true
			}
		} else if c > 1 {
			return true
		}
	}
	return false
}
