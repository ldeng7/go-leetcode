import "sort"

type Node struct {
	t int
	v string
}

type TimeMap struct {
	m map[string][]Node
}

func Constructor() TimeMap {
	return TimeMap{map[string][]Node{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Node{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	ar, ok := this.m[key]
	if !ok {
		return ""
	}
	i := sort.Search(len(ar), func(j int) bool {
		return ar[j].t > timestamp
	})
	if i != 0 {
		return ar[i-1].v
	}
	return ""
}
