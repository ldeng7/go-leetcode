func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

type SummaryRanges struct {
	ar [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{[][]int{}}
}

func (this *SummaryRanges) AddNum(val int) {
	itvl := []int{val, val}
	o, i := [][]int{}, 0
	for _, p := range this.ar {
		if itvl[1]+1 < p[0] {
			o = append(o, p)
		} else if itvl[0] > p[1]+1 {
			o = append(o, p)
			i++
		} else {
			itvl[0] = min(itvl[0], p[0])
			itvl[1] = max(itvl[1], p[1])
		}
	}
	if i != len(o) {
		o = append(o, nil)
		copy(o[i+1:], o[i:])
		o[i] = itvl
	} else {
		o = append(o, itvl)
	}
	this.ar = o
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.ar
}
