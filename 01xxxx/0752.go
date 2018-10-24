func openLock(deadends []string, target string) int {
	dm := map[string]bool{}
	for _, d := range deadends {
		dm[d] = true
	}
	if dm["0000"] {
		return -1
	}
	out := 0
	vm := map[string]bool{}
	q := []string{"0000"}
	for len(q) > 0 {
		out++
		for i := len(q); i > 0; i-- {
			s := q[len(q)-1]
			q = q[:len(q)-1]
			for j := 0; j < len(s); j++ {
				cal := func(k int8) bool {
					k = int8(s[j]-'0') + k
					if k < 0 {
						k += 10
					} else if k >= 10 {
						k -= 10
					}
					bs := []byte(s)
					bs[j] = '0' + byte(k)
					s1 := string(bs)
					if s1 == target {
						return true
					}
					if !dm[s1] && !vm[s1] {
						q = append(q, "")
						copy(q[1:], q[:len(q)-1])
						q[0] = s1
					}
					vm[s1] = true
					return false
				}
				if cal(-1) {
					return out
				}
				if cal(1) {
					return out
				}
			}
		}
	}
	return -1
}
