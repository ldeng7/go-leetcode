func addStrings(num1 string, num2 string) string {
	out := ""
	i, j := len(num1)-1, len(num2)-1
	var c byte = 0
	for i >= 0 || j >= 0 {
		var a, b byte = 0, 0
		if i >= 0 {
			a = num1[i] - '0'
			i--
		}
		if j >= 0 {
			b = num2[j] - '0'
			j--
		}
		s := a + b + c
		out = string('0'+s%10) + out
		c = s / 10
	}
	if c != 0 {
		return "1" + out
	}
	return out
}
