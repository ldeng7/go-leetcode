func findDiagonalOrder(nums [][]int) []int {
	c := 0
	for _, ar := range nums {
		c += len(ar)
	}
	o := make([]int, 0, c)
	q := make([]uint64, 1, c)
	l := uint64(len(nums) - 1)
	for 0 != len(q) {
		q0 := q[0]
		q = q[1:]
		i, j := q0>>17, q0&0x1ffff
		ar := nums[i]
		o = append(o, ar[j])
		if j == 0 && i < l {
			q = append(q, ((i+1)<<17)|j)
		}
		if j < uint64(len(ar))-1 {
			q = append(q, (i<<17)|(j+1))
		}
	}
	return o
}
