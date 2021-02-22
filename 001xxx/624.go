import "strings"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxLengthBetweenEqualCharacters(s string) int {
	o := -1
	for i := 0; i < len(s)/2; i++ {
		if j := strings.LastIndexByte(s, s[i]); j != -1 && j != i {
			o = max(o, j-i-1)
		}
	}
	return o
}
