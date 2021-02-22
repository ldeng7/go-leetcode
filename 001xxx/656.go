type OrderedStream struct {
	i  int
	ar []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{1, make([]string, n+1)}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.ar[id] = value
	if id == this.i {
		for id < len(this.ar) && 0 != len(this.ar[id]) {
			id++
		}
		o := this.ar[this.i:id]
		this.i = id
		return o
	} else {
		return []string{}
	}
}
