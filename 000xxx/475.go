import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	m := 0
	for _, o := range houses {
		i := sort.Search(len(heaters), func(j int) bool {
			return heaters[j] >= o
		})
		var d int
		if i != len(heaters) {
			d = heaters[i] - o
			if i != 0 && o-heaters[i-1] < d {
				d = o - heaters[i-1]
			}
		} else if i == len(heaters) {
			d = o - heaters[len(heaters)-1]
		}
		if d > m {
			m = d
		}
	}
	return m
}
