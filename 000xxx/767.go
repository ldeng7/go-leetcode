import "sort"

func reorganizeString(s string) string {
	cnt := make([]int, 26)
	out := []byte(s)
	for _, b := range out {
		cnt[b-'a'] += 100
	}
	for i := 0; i < 26; i++ {
		cnt[i] += i
	}
	sort.Ints(cnt)
	j := 1
	for _, c := range cnt {
		k := c / 100
		b := 'a' + byte(c%100)
		if k > (len(s)+1)/2 {
			return ""
		}
		for i := 0; i < k; i++ {
			if j >= len(s) {
				j = 0
			}
			out[j] = b
			j += 2
		}
	}
	return string(out)
}
