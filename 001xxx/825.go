type ocValType = [2]int

type OrderedContainer struct {
	arr []ocValType
	cnt int
	sum int
}

func (oc *OrderedContainer) init() *OrderedContainer {
	oc.arr = make([]ocValType, 0, 32)
	return oc
}

func (oc *OrderedContainer) min() int {
	return oc.arr[0][0]
}

func (oc *OrderedContainer) max() int {
	return oc.arr[len(oc.arr)-1][0]
}

func (oc *OrderedContainer) lowerBound(val int) int {
	i, j := 0, len(oc.arr)
	for i < j {
		h := i + (j-i)>>1
		if oc.arr[h][0] < val {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oc *OrderedContainer) add(val int) {
	i := oc.lowerBound(val)
	if i != len(oc.arr) {
		if oc.arr[i][0] != val {
			var e ocValType
			oc.arr = append(oc.arr, e)
			copy(oc.arr[i+1:], oc.arr[i:])
			oc.arr[i][0], oc.arr[i][1] = val, 1
		} else {
			oc.arr[i][1]++
		}
	} else {
		oc.arr = append(oc.arr, [2]int{val, 1})
	}
	oc.cnt++
	oc.sum += val
}

func (oc *OrderedContainer) remove(val int) {
	i := oc.lowerBound(val)
	if i == len(oc.arr) || oc.arr[i][0] != val {
		return
	}
	oc.arr[i][1]--
	if oc.arr[i][1] == 0 {
		if i != 0 {
			if i != len(oc.arr)-1 {
				copy(oc.arr[i:], oc.arr[i+1:])
			}
			oc.arr = oc.arr[:len(oc.arr)-1]
		} else {
			oc.arr = oc.arr[1:]
		}
	}
	oc.cnt--
	oc.sum -= val
}

type MKAverage struct {
	m, k    int
	q       []int
	p, l, r *OrderedContainer
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m, k, make([]int, 0, 32),
		(&OrderedContainer{}).init(),
		(&OrderedContainer{}).init(),
		(&OrderedContainer{}).init(),
	}
}

func (a *MKAverage) AddElement(num int) {
	if a.m == len(a.q) {
		a.pop()
	}
	a.push(num)
}

func (a *MKAverage) CalculateMKAverage() int {
	if a.m > len(a.q) {
		return -1
	}
	return a.p.sum / a.p.cnt
}

func (a *MKAverage) pop() {
	val := a.q[0]
	a.q = a.q[1:]
	if val >= a.p.min() && val <= a.p.max() {
		a.p.remove(val)
	} else if val <= a.l.max() {
		a.l.remove(val)
		val = a.p.min()
		a.l.add(val)
		a.p.remove(val)
	} else {
		a.r.remove(val)
		val = a.p.max()
		a.r.add(val)
		a.p.remove(val)
	}
}

func (a *MKAverage) push(val int) {
	a.q = append(a.q, val)
	if a.k > a.r.cnt {
		if a.k <= a.l.cnt {
			if t := a.l.max(); val < t {
				a.l.add(val)
				a.l.remove(t)
				val = t
			}
			a.r.add(val)
		} else {
			a.l.add(val)
		}
		return
	}
	if t := a.l.max(); val < t {
		a.l.add(val)
		a.l.remove(t)
		a.p.add(t)
	} else if t := a.r.min(); val > t {
		a.r.add(val)
		a.r.remove(t)
		a.p.add(t)
	} else {
		a.p.add(val)
	}
}
