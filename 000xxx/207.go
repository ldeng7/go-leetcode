func canFinish(numCourses int, prerequisites [][]int) bool {
	t := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		t[i] = []int{}
	}
	v := make([]int, numCourses)
	for _, p := range prerequisites {
		t[p[1]] = append(t[p[1]], p[0])
	}

	var cal func(int) bool
	cal = func(i int) bool {
		if v[i] == -1 {
			return false
		}
		if v[i] == 1 {
			return true
		}
		v[i] = -1
		for _, j := range t[i] {
			if !cal(j) {
				return false
			}
		}
		v[i] = 1
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !cal(i) {
			return false
		}
	}
	return true
}
