func countTriplets(A []int) int {
	t := [65536]int{}
	for _, i := range A {
		for _, j := range A {
			t[i&j]++
		}
	}
	o := 0
	for k, c := range t {
		if c == 0 {
			continue
		}
		for _, a := range A {
			if k&a == 0 {
				o += c
			}
		}
	}
	return o
}
