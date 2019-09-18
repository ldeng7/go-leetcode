import "sort"

type DinnerPlates struct {
	capa    int
	indices []int
	stacks  [][]int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{capacity, nil, nil}
}

func (this *DinnerPlates) Push(val int) {
	if 0 != len(this.indices) {
		i := this.indices[0]
		st := this.stacks[i]
		this.stacks[i] = append(st, val)
		if len(st)+1 == this.capa {
			this.indices = this.indices[1:]
		}
	} else {
		this.stacks = append(this.stacks, []int{val})
		if this.capa > 1 {
			this.indices = append(this.indices, len(this.stacks)-1)
		}
	}
}

func (this *DinnerPlates) Pop() int {
	return this.PopAtStack(len(this.stacks) - 1)
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(this.stacks) {
		return -1
	}
	st := this.stacks[index]
	l := len(st)
	if l == 0 {
		return -1
	}
	l--
	o := st[l]
	this.stacks[index] = st[:l]

	if l == this.capa-1 {
		if l == 0 && index == len(this.stacks)-1 {
			this.stacks = this.stacks[:len(this.stacks)-1]
		} else {
			i := sort.Search(len(this.indices), func(j int) bool {
				return this.indices[j] >= index
			})
			if i != len(this.indices) {
				this.indices = append(this.indices, 0)
				copy(this.indices[i+1:], this.indices[i:])
				this.indices[i] = index
			} else {
				this.indices = append(this.indices, index)
			}
		}
	} else if l == 0 && index == len(this.stacks)-1 {
		this.stacks = this.stacks[:len(this.stacks)-1]
		this.indices = this.indices[:len(this.indices)-1]
	}
	return o
}
