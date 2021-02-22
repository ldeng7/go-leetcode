func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	ds := binaryMatrix.Dimensions()
	o := ds[1]
	for i := 0; i < ds[0]; i++ {
		l, r := 0, ds[1]
		for l < r {
			m := l + (r-l)>>1
			if 0 != binaryMatrix.Get(i, m) {
				r = m
			} else {
				l = m + 1
			}
		}
		if l < o {
			o = l
		}
	}
	if o == ds[1] {
		return -1
	}
	return o
}
