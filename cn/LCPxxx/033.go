func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func cal(a, b int) int {
	if a == 0 {
		return 0
	} else if a%b == 0 {
		return a / b
	}
	return a/b + 1
}

func storeWater(bucket []int, vat []int) int {
	o := 0
	for i, v := range vat {
		if bucket[i] == 0 && v != 0 {
			o++
			bucket[i]++
		}
	}
	for {
		t, m := -1, 0
		for i, v := range vat {
			t = max(cal(v, bucket[i]), t)
		}
		ar := make([]int, 0, 16)
		for i, v := range vat {
			if t == cal(v, bucket[i]) {
				ar = append(ar, i)
			}
		}
		for _, i := range ar {
			bucket[i]++
		}
		for i, v := range vat {
			m = max(cal(v, bucket[i]), m)
		}
		m += len(ar)

		if t < m {
			o += t
			break
		} else {
			o += len(ar)
		}
	}
	return o
}
