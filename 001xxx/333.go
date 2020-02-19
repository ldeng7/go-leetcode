import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	ar := make([][2]int, 0, len(restaurants))
	for _, res := range restaurants {
		if veganFriendly == 1 && res[2] == 0 {
			continue
		}
		if res[3] <= maxPrice && res[4] <= maxDistance {
			ar = append(ar, [2]int{res[1], res[0]})
		}
	}
	sort.Slice(ar, func(i, j int) bool {
		a0, a1, b0, b1 := ar[i][0], ar[i][1], ar[j][0], ar[j][1]
		return a0 > b0 || (a0 == b0 && a1 > b1)
	})
	o := make([]int, len(ar))
	for i, p := range ar {
		o[i] = p[1]
	}
	return o
}
