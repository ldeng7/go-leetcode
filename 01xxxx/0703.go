import "sort"

type KthLargest struct {
	k    int
	heap []int
	l    int
}

func Constructor(k int, nums []int) KthLargest {
	obj := KthLargest{k, make([]int, k), 0}
	sort.IntSlice(nums).Sort()
	i := len(nums) - k
	if i < 0 {
		i = 0
	}
	copy(obj.heap, nums[i:])
	obj.l = len(nums) - i
	return obj
}

func (this *KthLargest) Add(val int) int {
	i := sort.Search(this.l, func(i int) bool {
		return this.heap[i] >= val
	})
	if this.l < this.k {
		if i < this.l {
			copy(this.heap[i+1:], this.heap[i:this.l])
		}
		this.heap[i] = val
		this.l++
	} else {
		if i == 1 {
			this.heap[0] = val
		} else if i > 1 {
			copy(this.heap, this.heap[1:i])
			this.heap[i-1] = val
		}
	}
	return this.heap[0]
}
