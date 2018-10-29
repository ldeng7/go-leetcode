func cs(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	m := map[int]bool{
		cs(p1, p2): true, cs(p1, p3): true, cs(p1, p4): true, cs(p2, p3): true, cs(p2, p4): true, cs(p3, p4): true,
	}
	return len(m) == 2 && !m[0]
}
