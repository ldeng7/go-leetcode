import "strconv"

func compress(chars []byte) int {
	k := 0
	for i, j := 0, 0; i < len(chars); i = j {
		for j < len(chars) && chars[i] == chars[j] {
			j++
		}
		chars[k] = chars[i]
		k++
		if j-i == 1 {
			continue
		}
		for _, c := range []byte(strconv.Itoa(j - i)) {
			chars[k] = c
			k++
		}
	}
	return k
}
