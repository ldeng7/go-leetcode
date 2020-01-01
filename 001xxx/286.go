type CombinationIterator struct {
	s  []byte
	c  int
	cl int
	b  bool
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	l, cl := len(characters), combinationLength
	s := make([]byte, l)
	for i := 0; i < l; i++ {
		s[l-i-1] = characters[i]
	}
	c := ((1 << cl) - 1) << (l - cl)
	return CombinationIterator{s, c, cl, true}
}

func (this *CombinationIterator) Next() string {
	o := make([]byte, 0, this.cl)
	for i := len(this.s) - 1; i >= 0; i-- {
		if this.c>>i&1 != 0 {
			o = append(o, this.s[i])
		}
	}
	a := this.c + 1
	b := this.c & a
	if b == 0 {
		this.b = false
		return string(o)
	}
	c := b & -b
	b = (b & (^c)) | (c >> 1)
	c >>= 2
	d := this.c & (^a)
	for d > 0 && d < c {
		d <<= 1
	}
	this.c = b | d
	return string(o)
}

func (this *CombinationIterator) HasNext() bool {
	return this.b
}
