func minFlips(target string) int {
	o, t := 0, byte(0)
	for i := 0; i < len(target); i++ {
		if c := target[i] - '0'; c != t {
			t = c
			o++
		}
	}
	return o
}
