func distributeCandies(candies int, numPeople int) []int {
	o := make([]int, numPeople)
	i, c := 0, 0
	for {
		candies -= c
		if candies <= i {
			o[i%numPeople] += candies
			break
		}
		c++
		o[i%numPeople] += c
		i++
	}
	return o
}
