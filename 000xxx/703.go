import "sort"

type SortedArray []int

func (sa SortedArray) LowerBound(bound int) int {
	return sort.Search(len(sa), func(j int) bool {
		return sa[j] >= bound
	})
}

func (sa SortedArray) Add(item int) SortedArray {
	i := sa.LowerBound(item)
	if i != len(sa) {
		sa = append(sa, 0)
		copy(sa[i+1:], sa[i:])
		sa[i] = item
	} else {
		sa = append(sa, item)
	}
	return sa
}

type KthLargest struct {
	k  int
	sa SortedArray
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	i := len(nums) - k
	if i < 0 {
		i = 0
	}
	l := len(nums) - i
	sa := make([]int, l)
	copy(sa, nums[i:])
	return KthLargest{k, sa}
}

func (this *KthLargest) Add(val int) int {
	if len(this.sa) == this.k {
		x := this.sa[0]
		if val <= x {
			return x
		}
		this.sa = this.sa[1:]
	}
	this.sa = this.sa.Add(val)
	return this.sa[0]
}
