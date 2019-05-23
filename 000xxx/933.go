import "sort"

type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.arr = append(this.arr, t)
	i := sort.Search(len(this.arr), func(j int) bool {
		return this.arr[j] >= t-3000
	})
	this.arr = this.arr[i:]
	return len(this.arr)
}
