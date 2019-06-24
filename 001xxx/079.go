func numTilePossibilities(tiles string) int {
	l := len(tiles)
	ar, m, bs := make([]bool, l), map[string]bool{}, []byte{}
	var cal func()
	cal = func() {
		if len(bs) > 0 {
			m[string(bs)] = true
		}
		if len(bs) == l {
			return
		}
		for i := 0; i < l; i++ {
			if ar[i] {
				continue
			}
			ar[i] = true
			bs = append(bs, tiles[i])
			cal()
			bs = bs[:len(bs)-1]
			ar[i] = false
		}
	}
	cal()
	return len(m)
}

