type Order struct {
	s string
	t int
}

type Stat struct {
	t, n int
}

type UndergroundSystem struct {
	orders map[int]Order
	m      map[string]*Stat
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{make(map[int]Order, 2048), make(map[string]*Stat, 2048)}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.orders[id] = Order{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	order := this.orders[id]
	s := order.s + ":" + stationName
	if st, ok := this.m[s]; ok {
		st.t += t - order.t
		st.n++
	} else {
		this.m[s] = &Stat{t - order.t, 1}
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	st := this.m[startStation+":"+endStation]
	return float64(st.t) / float64(st.n)
}
