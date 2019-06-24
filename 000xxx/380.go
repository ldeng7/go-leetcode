import "math/rand"

type RandomizedSet struct {
	ar []int
	m  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.ar = append(this.ar, val)
	this.m[val] = len(this.ar) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.m[val]
	if !ok {
		return false
	}
	l := len(this.ar)
	b := this.ar[l-1]
	this.ar[i], this.m[b] = b, i
	this.ar = this.ar[:l-1]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	l := len(this.ar)
	if 0 == l {
		return 0
	}
	return this.ar[rand.Int()%l]
}
