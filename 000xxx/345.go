func reverseVowels(s string) string {
	bs := []byte(s)
	ids := []int{}
	for i, b := range bs {
		switch b {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			ids = append(ids, i)
		}
	}
	for i := 0; i < len(ids)>>1; i++ {
		bs[ids[i]], bs[ids[len(ids)-i-1]] = bs[ids[len(ids)-i-1]], bs[ids[i]]
	}
	return string(bs)
}
