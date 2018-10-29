type NumArray struct {
	ar []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for _, n := range this.ar[i : j+1] {
		sum += n
	}
	return sum
}
