func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	l1, l2, l3 := len(arr1), len(arr2), len(arr3)
	l := l1
	if l2 > l {
		l = l2
	}
	if l3 > l {
		l = l3
	}
	o := make([]int, 0, l)
	for i, j, k := 0, 0, 0; i < l1 && j < l2 && k < l3; i, j, k = i+1, j+1, k+1 {
		a, b, c := arr1[i], arr2[j], arr3[k]
		if a == b && b == c {
			o = append(o, a)
		}
		if a > b || a > c {
			i--
		}
		if b > a || b > c {
			j--
		}
		if c > a || c > b {
			k--
		}
	}
	return o
}
