func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	t, c := tomatoSlices, cheeseSlices
	d := t - (c << 1)
	if (d < 0) || d&1 == 1 || d>>1 > c {
		return nil
	}
	a := d >> 1
	return []int{a, c - a}
}
