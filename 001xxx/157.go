import "sort"

type MajorityChecker struct {
	m [][]int
	f [][2]int
}

func Constructor(arr []int) MajorityChecker {
	ma := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > ma {
			ma = arr[i]
		}
	}
	m := make([][]int, ma)
	for i, n := range arr {
		m[n-1] = append(m[n-1], i)
	}

	f := [][2]int{}
	for i := 0; i < ma; i++ {
		if l := len(m[i]); 0 != l {
			f = append(f, [2]int{l, i})
		}
	}
	sort.Slice(f, func(i, j int) bool { return f[i][0] > f[j][0] })
	return MajorityChecker{m, f}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	for _, p := range this.f {
		if p[0] < threshold {
			return -1
		}
		ar := this.m[p[1]]
		i := sort.Search(len(ar), func(j int) bool { return ar[j] > right })
		j := sort.Search(len(ar), func(j int) bool { return ar[j] >= left })
		if i-j >= threshold {
			return p[1] + 1
		}
	}
	return -1
}
