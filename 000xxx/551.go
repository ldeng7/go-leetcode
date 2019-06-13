func checkRecord(s string) bool {
	ca, cl := 0, 0
	for _, c := range []byte(s) {
		switch c {
		case 'A':
			ca++
			if ca > 1 {
				return false
			}
			cl = 0
		case 'L':
			cl++
			if cl > 2 {
				return false
			}
		default:
			cl = 0
		}
	}
	return true
}
