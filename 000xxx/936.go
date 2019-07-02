import "bytes"

func movesToStamp(stamp string, target string) []int {
	ls, lt := len(stamp), len(target)
	bss, bst, t := []byte(stamp), []byte(target), make([]byte, lt)
	o, b := []int{}, true
	for b {
		if bytes.Compare(bst, t) == 0 {
			l := len(o)
			for i := 0; i < l>>1; i++ {
				o[i], o[l-1-i] = o[l-1-i], o[i]
			}
			return o
		}
		b = false
		for i := 0; i <= lt-ls; i++ {
			if bytes.Compare(t[i:i+ls], bst[i:i+ls]) == 0 {
				continue
			}
			b1 := true
			for j := 0; j < ls; j++ {
				if bss[j] != bst[i+j] && bst[i+j] != 0 {
					b1 = false
					break
				}
			}
			if b1 {
				o, b = append(o, i), true
				for j := i; j < i+ls; j++ {
					bst[j] = 0
				}
				b = true
				break
			}
		}
	}
	return []int{}
}
