import "math"

func getFactors(n int) [][]int {
	out := [][]int{}
	var cal func(int, int, []int)
	cal = func(n, j int, fs []int) {
		sq := int(math.Sqrt(float64(n)))
		for i := j; i <= sq; i++ {
			if n%i == 0 {
				fs1 := append(fs, i)
				r := n / i
				fsn := make([]int, len(fs1)+1)
				copy(fsn, fs1)
				fsn[len(fs1)] = r
				out = append(out, fsn)
				cal(r, i, fs1)
			}
		}
	}
	cal(n, 2, nil)
	return out
}
