import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func cal(ar []int) int {
	o, l := 0, len(ar)
	if l < 3 {
		for _, a := range ar {
			o += a / 3
		}
		return o
	}
	t := make([][5][3]int, l)
	for i := 0; i < l; i++ {
		for j := 0; j < 5; j++ {
			t[i][j][0], t[i][j][1], t[i][j][2] = -1, -1, -1
		}
	}
	for i, a0, a1 := 0, ar[0], ar[1]; i <= min(a1, 4); i++ {
		for j := 0; j <= min(a0, 2); j++ {
			t[1][i][j] = (a0-j)/3 + (a1-i)/3
		}
	}
	for i := 2; i < l; i++ {
		a, ap := ar[i], ar[i-1]
		js := min(a, 4)
		for j := js; j >= 0; j-- {
			ks := min(ap, 2)
			for k := ks; k >= 0; k-- {
				if j < js {
					t[i][j][k] = t[i][j+1][k]
				}
				if k < ks {
					t[i][j][k] = max(t[i][j][k], t[i][j][k+1])
				}
				for h := 0; h <= 2 && k+h <= min(ap, 4) && h+j <= a; h++ {
					t[i][j][k] = max(t[i][j][k], t[i-1][k+h][h]+(a-h-j)/3+h)
				}
			}
		}
	}
	return t[l-1][0][0]
}

func maxGroupNumber(tiles []int) int {
	sort.Ints(tiles)
	tiles = append(tiles, 1e10)
	o, p, c := 0, -1, 0
	ar := make([]int, 0, 64)
	for _, ti := range tiles {
		if ti == p {
			c++
			continue
		}
		ar = append(ar, c)
		if ti > p+1 {
			o += cal(ar)
			ar = make([]int, 0, 64)
		}
		p, c = ti, 1
	}
	return o
}
