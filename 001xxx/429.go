type FirstUnique struct {
	m map[int]int
	q []int
}

func Constructor(nums []int) FirstUnique {
	this := &FirstUnique{map[int]int{}, make([]int, 0, 16)}
	for _, n := range nums {
		this.Add(n)
	}
	return *this
}

func (this *FirstUnique) ShowFirstUnique() int {
	for 0 != len(this.q) {
		n := this.q[0]
		if this.m[n] > 1 {
			this.q = this.q[1:]
		} else {
			return n
		}
	}
	return -1
}

func (this *FirstUnique) Add(value int) {
	if c, ok := this.m[value]; ok {
		this.m[value] = c + 1
	} else {
		this.m[value] = 1
		this.q = append(this.q, value)
	}
}
