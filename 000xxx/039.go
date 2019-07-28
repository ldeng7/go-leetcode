import "sort"

func combinationSum(candidates []int, target int) [][]int {
	o, t := [][]int{}, []int{}
	sort.Ints(candidates)
	var cal func([]int, int)
	cal = func(cs []int, s int) {
		if s == target {
			t1 := make([]int, len(t))
			copy(t1, t)
			o = append(o, t1)
			return
		} else if s > target || 0 == len(cs) {
			return
		}
		for k, v := range cs {
			if s+v > target {
				break
			}
			t = append(t, v)
			cal(cs[k:], s+v)
			t = t[:len(t)-1]
		}
	}
	cal(candidates, 0)
	return o
}
