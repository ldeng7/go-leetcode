import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(people)))
	l, r, o := 0, len(people)-1, 0
	for l <= r {
		if people[l]+people[r] <= limit {
			r--
		}
		l++
		o++
	}
	return o
}
