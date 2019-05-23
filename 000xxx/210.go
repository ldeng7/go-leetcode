type Queue interface {
	Empty() bool
	Push(item int)
	Pop() (int, bool)
}

type ArrayQueue struct {
	arr []int
	i   int
}

func NewArrayQueue() Queue {
	return &ArrayQueue{[]int{}, 0}
}

func (q *ArrayQueue) Empty() bool {
	return len(q.arr)-q.i == 0
}

func (q *ArrayQueue) Push(item int) {
	if len(q.arr) <= 32 || q.i <= (len(q.arr)>>1) {
		q.arr = append(q.arr, item)
	} else {
		arr := make([]int, len(q.arr)-q.i+1)
		copy(arr, q.arr[q.i:])
		arr[len(arr)-1] = item
		q.arr = arr
		q.i = 0
	}
}

func (q *ArrayQueue) Pop() (int, bool) {
	if len(q.arr)-q.i != 0 {
		item := q.arr[q.i]
		q.i++
		return item, true
	}
	return 0, false
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	out := []int{}
	inArr := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		inArr[p[0]]++
	}
	q := NewArrayQueue()
	for i := 0; i < numCourses; i++ {
		if inArr[i] == 0 {
			q.Push(i)
		}
	}
	for !q.Empty() {
		j, _ := q.Pop()
		out = append(out, j)
		for _, k := range graph[j] {
			inArr[k]--
			if inArr[k] == 0 {
				q.Push(k)
			}
		}
	}
	if len(out) != numCourses {
		return []int{}
	}
	return out
}
