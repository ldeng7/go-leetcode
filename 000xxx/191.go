func hammingWeight(num uint32) int {
	var out uint32
	for num != 0 {
		out += num & 1
		num >>= 1
	}
	return int(out)
}
