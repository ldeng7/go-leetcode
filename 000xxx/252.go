import "sort"

type Pairs [][]int

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func canAttendMeetings(intervals [][]int) bool {
	var ms Pairs = intervals
	sort.Sort(ms)
	for i := 1; i < len(ms); i++ {
		if ms[i][0] < ms[i-1][1] {
			return false
		}
	}
	return true
}
