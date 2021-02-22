import "sort"

type SparseVector struct {
	ar []uint64
}

func Constructor(nums []int) SparseVector {
	ar := make([]uint64, 0, 32)
	for i, n := range nums {
		if n != 0 {
			ar = append(ar, (uint64(i)<<32)|uint64(n))
		}
	}
	return SparseVector{ar}
}

func (this *SparseVector) dotProduct(vec SparseVector) int {
	o := 0
	ar, ar1 := this.ar, vec.ar
	if len(ar) > len(ar1) {
		ar, ar1 = ar1, ar
	}
	for _, n := range ar {
		i, v := n>>32, n&0xff
		k := sort.Search(len(ar1), func(j int) bool {
			return ar1[j] >= i<<32
		})
		if k != len(ar1) {
			n1 := ar1[k]
			i1, v1 := n1>>32, n1&0xff
			if i1 == i {
				o += int(v * v1)
			}
		}
	}
	return o
}
