func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func checkPalindromeFormation(a string, b string) bool {
	i, j, l := 0, 0, len(a)
	for ; i < l && a[i] == b[l-1-i]; i++ {
	}
	for ; j < l && a[l-1-j] == b[j]; j++ {
	}
	la := (l - 1) / 2
	ra := l - 1 - la
	for ; la >= 0 && a[la] == a[ra]; la, ra = la-1, ra+1 {
	}
	lb := (l - 1) / 2
	rb := l - 1 - lb
	for ; lb >= 0 && b[lb] == b[rb]; lb, rb = lb-1, rb+1 {
	}
	c := max(i, j)*2 + max(ra-la-1, rb-lb-1)
	return c >= l
}

