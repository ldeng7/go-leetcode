import "sort"
import "math/rand"

type Solution struct {
	n int
	m map[int]int
	s []int
}

func Constructor(n int, bl []int) Solution {
	sort.IntSlice(bl).Sort()
	anc, pb := 0, -1
	s := Solution{n - len(bl), map[int]int{anc: 0}, []int{anc}}
	for i, b := range bl {
		if b == pb+1 {
			s.m[anc]++
		} else {
			anc += b - pb - 1
			s.m[anc] = i + 1
			s.s = append(s.s, anc)
		}
		pb = b
	}
	return s
}

func (this *Solution) Pick() int {
	j := int(rand.Int31n(int32(this.n)))
	k := sort.Search(len(this.s), func(i int) bool {
		return this.s[i] > j
	})
	return j + this.m[this.s[k-1]]
}
