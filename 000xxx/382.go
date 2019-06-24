import "math/rand"

type Solution struct {
	h *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	o, n, i := this.h.Val, this.h.Next, 2
	for nil != n {
		j := rand.Int() % i
		if j == 0 {
			o = n.Val
		}
		i++
		n = n.Next
	}
	return o
}
