type NumArray struct {
	l int
	t []int
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	t := make([]int, l<<1)
	for i := l; i < l<<1; i++ {
		t[i] = nums[i-l]
	}
	for i := l - 1; i > 0; i-- {
		t[i] = t[i*2] + t[i*2+1]
	}
	return NumArray{l, t}
}

func (this *NumArray) Update(i int, val int) {
	i += this.l
	this.t[i] = val
	for i > 0 {
		this.t[i>>1] = this.t[i] + this.t[i^1]
		i >>= 1
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	s := 0
	i, j = i+this.l, j+this.l
	for ; i <= j; i, j = i>>1, j>>1 {
		if (i & 1) == 1 {
			s, i = s+this.t[i], i+1
		}
		if (j & 1) == 0 {
			s, j = s+this.t[j], j-1
		}
	}
	return s
}
