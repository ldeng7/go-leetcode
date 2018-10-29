import "sort"

type MyCalendar struct {
	sl []int
	m  map[int]int
}

func Constructor() MyCalendar {
	return MyCalendar{[]int{}, map[int]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	j := sort.Search(len(this.sl), func(i int) bool {
		return this.sl[i] >= start
	})
	if (j != len(this.sl) && (this.sl[j] == start || end > this.sl[j])) || (j != 0 && this.m[this.sl[j-1]] > start) {
		return false
	}
	this.sl = append(this.sl, 0)
	if j != len(this.sl)-1 {
		copy(this.sl[j+1:], this.sl[j:])
	}
	this.sl[j] = start
	this.m[start] = end
	return true
}
