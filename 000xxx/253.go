import "sort"

func minMeetingRooms(intervals [][]int) int {
	bs, es := make([]int, len(intervals)), make([]int, len(intervals))
	for i, m := range intervals {
		bs[i], es[i] = m[0], m[1]
	}
	sort.Ints(bs)
	sort.Ints(es)
	c, ie := 0, 0
	for _, b := range bs {
		if b < es[ie] {
			c++
		} else {
			ie++
		}
	}
	return c
}
