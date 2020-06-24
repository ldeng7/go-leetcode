func largestMultipleOfThree(digits []int) string {
	ar, sum := [10]int{}, 0
	for _, d := range digits {
		ar[d]++
		sum += d
	}
	sum %= 3
	b := false
	if 0 != sum {
		for i := sum; i <= sum+6; i += 3 {
			if ar[i] > 0 {
				b = true
				ar[i]--
				break
			}
		}
		if !b {
			c := 2
			for i := 3 - sum; i <= 9-sum; i += 3 {
				if a := ar[i]; a > 0 {
					if a < c {
						c, ar[i] = c-a, 0
					} else {
						c, ar[i] = 0, a-c
					}
				}
			}
			b = 0 == c
		}
	} else {
		b = true
	}

	if !b {
		return ""
	}
	bs := make([]byte, 0, 32)
	for i := 9; i >= 0; i-- {
		for j := 0; j < ar[i]; j++ {
			bs = append(bs, '0'+byte(i))
		}
	}
	if len(bs) > 1 && bs[0] == '0' && bs[len(bs)-1] == '0' {
		return "0"
	}
	return string(bs)
}
