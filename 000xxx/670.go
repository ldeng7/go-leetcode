import "strconv"

func maximumSwap(num int) int {
	bs := []byte(strconv.Itoa(num))
	l := len(bs)
	out := make([]byte, l)
	copy(out, bs)
	for i := l - 2; i >= 0; i-- {
		if bs[i+1] > bs[i] {
			bs[i] = bs[i+1]
		}
	}
	for i := 0; i < l; i++ {
		if out[i] == bs[i] {
			continue
		}
		for j := l - 1; j > i; j-- {
			if out[j] == bs[i] {
				out[i], out[j] = out[j], out[i]
				k, _ := strconv.Atoi(string(out))
				return k
			}
		}
	}
	k, _ := strconv.Atoi(string(out))
	return k
}
