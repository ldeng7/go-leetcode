func largestOddNumber(num string) string {
	i := len(num) - 1
	for i >= 0 {
		if (num[i]-'0')&1 == 0 {
			i--
		} else {
			break
		}
	}
	return num[:i+1]
}
