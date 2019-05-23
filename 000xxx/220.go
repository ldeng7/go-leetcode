import "sort"

type SortedArray []int

func (sa SortedArray) LowerBound(bound int) int {
	return sort.Search(len(sa), func(j int) bool {
		return sa[j] >= bound
	})
}

func (sa SortedArray) Index(item int) int {
	i := sa.LowerBound(item)
	if i == len(sa) || sa[i] != item {
		return -1
	}
	return i
}

func (sa SortedArray) Add(item int) SortedArray {
	i := sa.LowerBound(item)
	if i != len(sa) {
		sa = append(sa, 0)
		copy(sa[i+1:], sa[i:])
		sa[i] = item
	} else {
		sa = append(sa, item)
	}
	return sa
}

func (sa SortedArray) Remove(item int) SortedArray {
	i := sa.Index(item)
	if -1 == i {
		return sa
	}
	if i != len(sa)-1 {
		copy(sa[i:], sa[i+1:])
	}
	return sa[:len(sa)-1]
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	sa := SortedArray{}
	j := 0
	for i, num := range nums {
		if i-j > k {
			sa = sa.Remove(nums[j])
			j++
		}
		b := sa.LowerBound(num - t)
		if b != len(sa) && sa[b] <= num+t {
			return true
		}
		sa = sa.Add(num)
	}
	return false
}
