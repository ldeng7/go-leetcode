func checkIfExist(arr []int) bool {
	m := map[int]bool{}
	z := 0
	for _, a := range arr {
		if 0 != a {
			m[a] = true
		} else {
			z++
		}
	}
	if z >= 2 {
		return true
	}
	for _, a := range arr {
		if m[a<<1] || (a&1 == 0 && m[a>>1]) {
			println(a)
			return true
		}
	}
	return false
}
