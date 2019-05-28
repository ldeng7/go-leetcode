import "strings"

func repeatedStringMatch(A string, B string) int {
	s := A
	for i := 1; i <= len(B)/len(A)+2; i++ {
		if strings.Contains(s, B) {
			return i
		}
		s += A
	}
	return -1
}
