import "math"

func racecar(target int) int {
	t := make([]int, target+1)
	for i := 1; i <= target; i++ {
		t[i] = math.MaxInt64
		j, c := 1, 1
		for ; j < i; j = (1 << uint(c)) - 1 {
			for k, c1 := 0, 0; k < j; k = (1 << uint(c1)) - 1 {
				t1 := c + c1 + 2 + t[i-j+k]
				if t1 < t[i] {
					t[i] = t1
				}
				c1++
			}
			c++
		}
		t1 := c
		if i != j {
			t1 += 1 + t[j-i]
		}
		if t1 < t[i] {
			t[i] = t1
		}
	}
	return t[target]
}
