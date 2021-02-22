func breakfastNumber(staple []int, drinks []int, x int) int {
	ar := make([]int, x+1)
	for _, s := range staple {
		if s < x {
			ar[s]++
		}
	}
	for i := 2; i < x; i++ {
		ar[i] += ar[i-1]
	}
	o := 0
	for _, d := range drinks {
		i := x - d
		if i <= 0 {
			continue
		}
		o = (o + ar[i]) % 1000000007
	}
	return o
}
