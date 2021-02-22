func minPartitions(n string) int {
	o := 0
	for i := 0; i < len(n); i++ {
		if v := int(n[i] - '0'); v > o {
			o = v
		}
		if o == 9 {
			return 9
		}
	}
	return o
}
