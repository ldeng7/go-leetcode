import "sort"

func halfQuestions(questions []int) int {
	cs := map[int]int{}
	for _, a := range questions {
		cs[a]++
	}
	ar := make([]int, 0, len(cs))
	for _, c := range cs {
		ar = append(ar, c)
	}
	sort.Ints(ar)
	o := 0
	for i, n := len(ar)-1, len(questions)/2; n > 0; i, n = i-1, n-ar[i] {
		o++
	}
	return o
}
