func findTheDifference(s string, t string) byte {
	var a byte = 0
	for _, c := range []byte(t) {
		a += c
	}
	for _, c := range []byte(s) {
		a -= c
	}
	return a
}
