type NumArray struct {
	t []int
}

func Constructor(nums []int) NumArray {
	t := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		t[i] = t[i-1] + nums[i-1]
	}
	return NumArray{t}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.t[j+1] - this.t[i]
}
