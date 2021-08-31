import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

type node struct {
	i int8
	j int8
	n int32
}

func maxScore(nums []int) int {
	o, l := 0, len(nums)
	m := map[int8]bool{}
	ar := make([]node, 0, l*l)
	for i := 0; i < l; i++ {
		n := nums[i]
		for j := i + 1; j < l; j++ {
			ar = append(ar, node{int8(i), int8(j), int32(gcd(n, nums[j]))})
		}
	}
	sort.Slice(ar, func(i, j int) bool { return ar[i].n > ar[j].n })

	var cal func(int, int, int)
	cal = func(i, c, k int) {
		if k == 0 {
			o = max(o, c)
			return
		} else if i < len(ar) && (k+1)*k/2*int(ar[i].n)+c <= o {
			return
		}
		for j := i; j < len(ar); j++ {
			n := ar[j]
			if m[n.i] || m[n.j] {
				continue
			}
			m[n.i], m[n.j] = true, true
			cal(j+1, c+int(n.n)*k, k-1)
			m[n.i], m[n.j] = false, false
		}
	}

	cal(0, 0, l/2)
	return o
}
