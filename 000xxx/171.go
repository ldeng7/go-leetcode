func titleToNumber(s string) int {
	sum := 0
	base := 1
	for i := len(s) - 1; i >= 0; i-- {
		sum += int(s[i]-'A'+1) * base
		base *= 26
	}
	return sum
}
