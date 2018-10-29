var m = map[int]bool{6: true, 28: true, 496: true, 8128: true, 33550336: true}

func checkPerfectNumber(num int) bool {
	return m[num]
}
