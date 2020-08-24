func numOfSubarrays(arr []int) int {
	o, p := 0, [2]int{}
	for _, a := range arr {
		a &= 1
		if a == 1 {
			p[0], p[1] = p[1], p[0]
		}
		p[a] = (p[a] + 1) % 1000000007
		o = (o + p[1]) % 1000000007
	}
	return o
}
