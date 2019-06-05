type StringIterator struct {
	bs []byte
	cs []int
	i  int
}

func Constructor(compressedString string) StringIterator {
	bs, cs := []byte{}, []int{}
	s := []byte(compressedString)
	i := 0
	for i < len(s) {
		a := s[i]
		i++
		n := 0
		for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			n *= 10
			n += int(s[i] - '0')
		}
		bs = append(bs, a)
		cs = append(cs, n)
	}
	return StringIterator{bs, cs, 0}
}

func (this *StringIterator) Next() byte {
	i, l := this.i, len(this.bs)
	if i != l {
		c := this.cs[i]
		if c > 1 {
			this.cs[i]--
		} else {
			this.i++
		}
		return this.bs[i]
	}
	return ' '
}

func (this *StringIterator) HasNext() bool {
	return this.i < len(this.bs)
}
