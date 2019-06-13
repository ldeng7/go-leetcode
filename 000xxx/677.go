type MapSum struct {
	m  map[string]int
	md map[string]int
}

func Constructor() MapSum {
	return MapSum{map[string]int{}, map[string]int{}}
}

func (this *MapSum) Insert(key string, val int) {
	d := val - this.m[key]
	this.m[key] = val
	for i := 1; i < len(key); i++ {
		this.md[key[:i]] += d
	}
}

func (this *MapSum) Sum(prefix string) int {
	return this.m[prefix] + this.md[prefix]
}
