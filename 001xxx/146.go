import "sort"

type SnapshotArray struct {
	ars [][2][]int
	cnt int
}

func Constructor(length int) SnapshotArray {
	ars := make([][2][]int, length)
	for i := 0; i < length; i++ {
		ars[i] = [2][]int{[]int{0}, []int{0}}
	}
	return SnapshotArray{ars, 0}
}

func (this *SnapshotArray) Set(index int, val int) {
	pair := &(this.ars[index])
	(*pair)[0], (*pair)[1] = append((*pair)[0], this.cnt), append((*pair)[1], val)
}

func (this *SnapshotArray) Snap() int {
	this.cnt++
	return this.cnt - 1
}

func (this *SnapshotArray) Get(index int, snapId int) int {
	pair := this.ars[index]
	cs := pair[0]
	i := sort.Search(len(cs), func(j int) bool { return cs[j] > snapId })
	return pair[1][i-1]
}
