func strobogrammaticInRange(low string, high string) int {
	var c int
	var cal func(string)
	cal = func(s string) {
		if len(s) >= len(low) && len(s) <= len(high) {
			if len(s) == len(high) && s > high {
				return
			}
			if (len(s) <= 1 || s[0] != '0') && (len(s) > len(low) || s >= low) {
				c++
			}
		}
		if len(s) > len(high)-2 {
			return
		}
		cal("0" + s + "0")
		cal("1" + s + "1")
		cal("6" + s + "9")
		cal("8" + s + "8")
		cal("9" + s + "6")
	}
	cal("")
	cal("0")
	cal("1")
	cal("8")
	return c
}
