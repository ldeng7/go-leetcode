import (
	"sort"
	"strconv"
	"strings"
)

type Tran struct {
	index  int
	time   int64
	amount int64
	city   string
}

func invalidTransactions(transactions []string) []string {
	m := map[string][]*Tran{}
	for i, s := range transactions {
		es := strings.Split(s, ",")
		if len(es) >= 4 {
			name := es[0]
			time, _ := strconv.ParseInt(es[1], 10, 16)
			amount, _ := strconv.ParseInt(es[2], 10, 16)
			tran := &Tran{i, time, amount, es[3]}
			m[name] = append(m[name], tran)
		}
	}

	s := map[int]bool{}
	for _, ar := range m {
		sort.Slice(ar, func(i, j int) bool {
			return ar[i].time < ar[j].time
		})
		if ar[0].amount > 1000 {
			s[ar[0].index] = true
		}
		for i := 1; i < len(ar); i++ {
			t := ar[i]
			if t.amount > 1000 {
				s[t.index] = true
			}
			for j := i - 1; j >= 0; j-- {
				t1 := ar[j]
				if t.time-t1.time > 60 {
					break
				}
				if t.city != t1.city {
					s[t.index], s[t1.index] = true, true
				}
			}
		}
	}

	o := make([]string, 0, len(transactions))
	for index, _ := range s {
		o = append(o, transactions[index])
	}
	return o
}
