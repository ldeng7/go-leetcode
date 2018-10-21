type MyHashMap struct {
	m map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{map[int]int{}}
}

func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

func (this *MyHashMap) Get(key int) int {
	v, ok := this.m[key]
	if !ok {
		return -1
	}
	return v
}

func (this *MyHashMap) Remove(key int) {
	delete(this.m, key)
}
