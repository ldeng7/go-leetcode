func arrayPairSum(nums []int) int {
	cs := [20001]int{}
	for _, n := range nums {
		cs[n+10000]++
	}
	o, d := 0, 1
	for i := 0; i <= 20000; i++ {
		if c := cs[i]; c != 0 {
			o += (i - 10000) * ((c + d) >> 1)
			d = (d ^ c) & 1
		}
	}
	return o
}
