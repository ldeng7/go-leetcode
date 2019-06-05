func rotateString(A string, B string) bool {
	l, lb := len(A), len(B)
	if l != lb {
		return false
	}
	if 0 == l {
		return true
	}
	for i := 0; i < l; i++ {
		if A[i] != B[0] {
			continue
		}
		if A[i:] == B[:l-i] && A[:i] == B[l-i:] {
			return true
		}
	}
	return false
}
