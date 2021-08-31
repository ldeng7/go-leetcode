func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func getCoprimes(nums []int, edges [][]int) []int {
	m := map[int][]int{}
	mp := map[int][][2]int{}
	s := map[int]bool{}
	for _, n := range nums {
		s[n] = true
	}
	for n1, _ := range s {
		for n2, _ := range s {
			if gcd(n1, n2) == 1 {
				m[n1] = append(m[n1], n2)
			}
		}
	}

	l := len(nums)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, 0, 16)
	}
	for _, e := range edges {
		a, b := e[0], e[1]
		t[a], t[b] = append(t[a], b), append(t[b], a)
	}

	o := make([]int, l)
	for i := 0; i < l; i++ {
		o[i] = -1
	}
	var cal func(int, int, int)
	cal = func(i, p, d int) {
		n, dm := nums[i], -1
		for _, n1 := range m[n] {
			ps := mp[n1]
			if l := len(ps); 0 != l {
				p := ps[l-1]
				if d := p[0]; d > dm {
					dm, o[i] = d, p[1]
				}
			}
		}
		mp[n] = append(mp[n], [2]int{d, i})
		for _, j := range t[i] {
			if j != p {
				cal(j, i, d+1)
			}
		}
		ps := mp[n]
		mp[n] = ps[:len(ps)-1]
	}
	cal(0, 0, 0)
	return o
}
