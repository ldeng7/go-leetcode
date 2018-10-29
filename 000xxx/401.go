import "fmt"

func readBinaryWatch(num int) []string {
	hs := [][]int{{0}, {1, 2, 4, 8}, {3, 5, 9, 6, 10}, {7, 11}}
	ms := [][]int{
		{0}, {1, 2, 4, 8, 16, 32}, {3, 5, 9, 17, 33, 6, 10, 18, 34, 12, 20, 36, 24, 40, 48},
		{7, 11, 19, 35, 13, 21, 37, 25, 41, 49, 14, 22, 38, 26, 42, 50, 28, 44, 52, 56},
		{15, 23, 39, 27, 43, 51, 29, 45, 53, 57, 30, 46, 54, 58}, {31, 47, 55, 59},
	}
	out := []string{}
	for i := 0; i <= num; i++ {
		j := num - i
		if i > 3 || j > 5 {
			continue
		}
		for _, h := range hs[i] {
			for _, m := range ms[j] {
				out = append(out, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return out
}
