import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l, r, o := 0, len(people)-1, 0
	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		r, o = r-1, o+1
	}
	return o
}
