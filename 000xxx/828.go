func uniqueLetterString(S string) int {
	pre, cur := [26]int{}, [26]int{}
	for i := 0; i < 26; i++ {
		pre[i], cur[i] = -1, -1
	}
	out, i := 0, 0
	for ; i < len(S); i++ {
		j := S[i] - 'A'
		p, pp := cur[j], pre[j]
		pre[j], cur[j] = p, i
		out += (i - p) * (p - pp)
	}
	for j := 0; j < 26; j++ {
		out += (i - cur[j]) * (cur[j] - pre[j])
	}
	return out
}
