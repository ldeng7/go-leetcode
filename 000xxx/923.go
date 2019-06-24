func threeSumMulti(A []int, target int) int {
	o := 0
	cs := [101]int{}
	for _, a := range A {
		cs[a]++
	}
	for i, ci := range cs {
		if ci == 0 {
			continue
		}
		for j := i; j <= 100; j++ {
			cj := cs[j]
			if cj == 0 {
				continue
			}
			k := target - i - j
			if k < j {
				break
			} else if k > 100 || cs[k] == 0 {
				continue
			}
			ck := cs[k]
			var o1 int
			if i == j {
				if j == k {
					o1 = ci * (ci - 1) * (ci - 2) / 6
				} else {
					o1 = ci * (ci - 1) * ck / 2
				}
			} else {
				if j == k {
					o1 = ci * cj * (cj - 1) / 2
				} else {
					o1 = ci * cj * ck
				}
			}
			o = (o + o1) % 1000000007
		}
	}
	return o
}
