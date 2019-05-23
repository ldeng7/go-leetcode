func reverseBits(num uint32) uint32 {
	var out uint32
	for i := uint(0); i < 32; i++ {
		out |= ((num >> i) & 1) << (31 - i)
	}
	return out
}
