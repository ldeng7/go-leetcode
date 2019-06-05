func maxSatisfied(customers []int, grumpy []int, X int) int {
	o1 := 0
	for i, c := range customers {
		if 0 == grumpy[i] {
			o1, customers[i] = o1+c, 0
		}
	}
	i, t := 0, 0
	for ; i < len(customers) && i < X; i++ {
		t += customers[i]
	}
	o := t
	for ; i < len(customers); i++ {
		t += customers[i] - customers[i-X]
		if t > o {
			o = t
		}
	}
	return o + o1
}
