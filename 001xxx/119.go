var m = [26]bool{
	true, false, false, false, true, false, false, false, true, false, false, false, false, false,
	true, false, false, false, false, false, true, false, false, false, false, false,
}

func removeVowels(S string) string {
	bs, i := make([]byte, len(S)), 0
	for _, c := range []byte(S) {
		if !m[c-'a'] {
			bs[i], i = c, i+1
		}
	}
	return string(bs[:i])
}
