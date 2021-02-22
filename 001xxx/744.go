func canEat(candiesCount []int, queries [][]int) []bool {
	ar := candiesCount
	for i := 1; i < len(ar); i++ {
		ar[i] += ar[i-1]
	}
	o := make([]bool, len(queries))
	for i, q := range queries {
		j, d := q[0], q[1]+1
		if j != 0 {
			j = ar[j-1]
		}
		o[i] = d*q[2] > j && ar[q[0]] >= d
	}
	return o
}
