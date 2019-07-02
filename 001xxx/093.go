func sampleStats(count []int) []float64 {
	min, max, sum, cnt, mode, cm := -1, 0, 0, 0, 0, 0
	for i, c := range count {
		if c == 0 {
			continue
		}
		if min == -1 {
			min = i
		}
		max = i
		sum += c * i
		cnt += c
		if c > cm {
			mode, cm = i, c
		}
	}
	var med float64
	if cnt&1 == 0 {
		n := (cnt >> 1) - 1
		for i, c := range count {
			if c == 0 {
				continue
			}
			n -= c
			if n < 0 {
				if n != -1 {
					med = float64(i)
				} else {
					for j := i + 1; ; j++ {
						if count[j] != 0 {
							med = float64(i+j) / 2
							break
						}
					}
				}
				break
			}
		}
	} else {
		n := cnt >> 1
		for i, c := range count {
			if c == 0 {
				continue
			}
			n -= c
			if n < 0 {
				med = float64(i)
				break
			}
		}
	}
	return []float64{
		float64(min),
		float64(max),
		float64(sum) / float64(cnt),
		med,
		float64(mode),
	}
}
