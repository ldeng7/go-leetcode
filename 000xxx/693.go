func hasAlternatingBits(n int) bool {
	or := n | n>>1
	return (n&(n>>1)) == 0 && (or&(or+1)) == 0
}
