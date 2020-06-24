func kidsWithCandies(candies []int, extraCandies int) []bool {
	m := 0
	for _, n := range candies {
		if n > m {
			m = n
		}
	}
	m -= extraCandies
	o := make([]bool, len(candies))
	for i, n := range candies {
		if n >= m {
			o[i] = true
		}
	}
	return o
}
