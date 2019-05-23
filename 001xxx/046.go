import "sort"

type SortedArray []int

func (sa SortedArray) Add(item int) SortedArray {
	i := sort.Search(len(sa), func(j int) bool {
		return sa[j] >= item
	})
	if i != len(sa) {
		sa = append(sa, 0)
		copy(sa[i+1:], sa[i:])
		sa[i] = item
	} else {
		sa = append(sa, item)
	}
	return sa
}

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	var ss SortedArray = stones
	for len(ss) > 1 {
		d := ss[len(ss)-1] - ss[len(ss)-2]
		ss = ss[:len(ss)-2]
		if 0 != d {
			ss = ss.Add(d)
		}
	}
	if 1 == len(ss) {
		return ss[0]
	}
	return 0
}
