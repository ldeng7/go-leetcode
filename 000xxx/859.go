func buddyStrings(A string, B string) bool {
	if A == B {
		c := [26]int{}
		for _, a := range A {
			c[a-'a']++
			if c[a-'a'] >= 2 {
				return true
			}
		}
		return false
	}

	if len(A) != len(B) {
		return false
	}
	var ca, cb byte
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			if ca == 1 {
				return false
			} else if ca == 0 {
				ca, cb = A[i], B[i]
			} else if ca != B[i] || cb != A[i] {
				return false
			} else {
				ca = 1
			}
		}
	}
	return ca == 1
}
