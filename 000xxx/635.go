type LogSystem struct {
	ids []int
	ts  []string
}

func Constructor() LogSystem {
	return LogSystem{[]int{}, []string{}}
}

func (this *LogSystem) Put(id int, timestamp string) {
	this.ids = append(this.ids, id)
	this.ts = append(this.ts, timestamp)
}

var m map[string]int = map[string]int{
	"Year": 4, "Month": 7, "Day": 10, "Hour": 13, "Minute": 16, "Second": 19,
}

func (this *LogSystem) Retrieve(s string, e string, gra string) []int {
	out := []int{}
	j := m[gra]
	for i, id := range this.ids {
		t := this.ts[i][:j]
		if t >= s[:j] && t <= e[:j] {
			out = append(out, id)
		}
	}
	return out
}
