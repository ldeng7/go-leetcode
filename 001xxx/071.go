func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) != len(str2) {
		if len(str1) < len(str2) {
			str1, str2 = str2, str1
		}
		for len(str1) > len(str2) {
			if str1[:len(str2)] != str2 {
				return ""
			}
			str1 = str1[len(str2):]
		}
		return gcdOfStrings(str1, str2)
	} else {
		if str1 == str2 {
			return str1
		} else {
			return ""
		}
	}
}
