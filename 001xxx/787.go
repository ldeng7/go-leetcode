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

func minChanges(nums []int, k int) int {
	fs := make([]map[int]int, k)
	for i := 0; i < k; i++ {
		fs[i] = map[int]int{}
	}
	mas, mass := make([]int, k), make([]int, k)
	for i, n := range nums {
		fs[i%k][n]++
	}

	for i, f := range fs {
		for _, c := range f {
			mas[i] = max(mas[i], c)
		}
	}
	mass[k-1] = mas[k-1]
	for i := k - 2; i >= 0; i-- {
		mass[i] = mass[i+1] + mas[i]
	}

	m := mas[0]
	for i := 1; i < k; i++ {
		m = min(m, mas[i])
	}
	o := mass[0] - m

	var cal func(int, int, int)
	cal = func(i, s, c int) {
		if i == k {
			if s == 0 {
				o = max(o, c)
			}
		} else if c+mass[i] > o {
			for j, f := range fs[i] {
				cal(i+1, s^j, c+f)
			}
		}
	}

	cal(0, 0, 0)
	return len(nums) - o
}
