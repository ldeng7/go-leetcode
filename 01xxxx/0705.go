type MyHashSet struct {
	m map[int]bool
}

func Constructor() MyHashSet {
	return MyHashSet{map[int]bool{}}
}

func (this *MyHashSet) Add(key int) {
	this.m[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.m[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.m[key]
}