func reverseSlice(ar []int) {
	e, h := len(ar)-1, len(ar)>>1
	for i := 0; i < h; i++ {
		ar[i], ar[e-i] = ar[e-i], ar[i]
	}
}

func addNegabinary(arr1 []int, arr2 []int) []int {
	l1, l2 := len(arr1), len(arr2)
	reverseSlice(arr1)
	reverseSlice(arr2)
	o, c := []int{}, 0
	for i := 0; i < l1 || i < l2 || c != 0; i++ {
		a, b := 0, 0
		if i < l1 {
			a = arr1[i]
		}
		if i < l2 {
			b = arr2[i]
		}
		a = a + b + c
		b = (a + 2) & 1
		o, c = append(o, b), (a-b)/-2
	}
	for len(o) > 1 && o[len(o)-1] == 0 {
		o = o[:len(o)-1]
	}
	reverseSlice(o)
	return o
}
