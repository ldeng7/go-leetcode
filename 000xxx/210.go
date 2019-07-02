func findOrder(numCourses int, prerequisites [][]int) []int {
	out := []int{}
	arr := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		arr[p[0]]++
	}
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if arr[i] == 0 {
			q = append(q, i)
		}
	}
	for 0 != len(q) {
		i := q[0]
		q = q[1:]
		out = append(out, i)
		for _, j := range graph[i] {
			arr[j]--
			if arr[j] == 0 {
				q = append(q, j)
			}
		}
	}
	if len(out) != numCourses {
		return []int{}
	}
	return out
}
