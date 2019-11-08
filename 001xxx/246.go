import "math"

func minimumMoves(arr []int) int {
	l := len(arr)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		ar := make([]int, l)
		for j := 0; j < l; j++ {
			ar[j] = math.MaxInt64
		}
		t[i] = ar
	}

	var cal func(int, int) int
	cal = func(l, r int) int {
		if l > r {
			return 0
		} else if l == r {
			t[l][r] = 1
		} else if l+1 == r {
			if arr[l] == arr[r] {
				t[l][r] = 1
			} else {
				t[l][r] = 2
			}
		} else if t[l][r] != math.MaxInt64 {
			return t[l][r]
		} else {
			t[l][r] = cal(l+1, r) + 1
			for i := l + 1; i <= r; i++ {
				if arr[i] == arr[l] {
					o := cal(l+1, i-1) + cal(i+1, r)
					if i == l+1 {
						o++
					}
					if o < t[l][r] {
						t[l][r] = o
					}
				}
			}
		}
		return t[l][r]
	}
	return cal(0, l-1)
}
