import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	o, t := [][]int{}, []int{}
	sort.Ints(candidates)
	var cal func([]int, int)
	cal = func(cs []int, tar int) {
		if 0 == tar {
			t1 := make([]int, len(t))
			copy(t1, t)
			o = append(o, t1)
			return
		} else if len(cs) == 0 || tar < cs[0] {
			return
		}
		t = append(t, cs[0])
		cal(cs[1:], tar-cs[0])
		t = t[:len(t)-1]
		i := 0
		for ; i+1 < len(cs) && cs[i] == cs[i+1]; i++ {
		}
		cal(cs[i+1:], tar)
	}
	cal(candidates, target)
	return o
}
