import "sort"

type oaElemType = int
type oaElemCmpCb = func(oaElemType, oaElemType) bool

type OrderedArray struct {
	arr    []oaElemType
	lessCb oaElemCmpCb
	eqCb   oaElemCmpCb
}

func (oa *OrderedArray) Init(arr []oaElemType, lessCb, eqCb oaElemCmpCb) *OrderedArray {
	oa.arr = arr
	if len(arr) > 1 {
		sort.Slice(arr, func(i, j int) bool { return lessCb(arr[i], arr[j]) })
	}
	oa.lessCb = lessCb
	oa.eqCb = eqCb
	return oa
}

func (oa *OrderedArray) LowerBound(item oaElemType) int {
	i, j := 0, len(oa.arr)
	for i < j {
		h := i + (j-i)>>1
		if oa.lessCb(oa.arr[h], item) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func (oa *OrderedArray) Add(item oaElemType) {
	i := oa.LowerBound(item)
	if i != len(oa.arr) {
		oa.arr = append(oa.arr, 0)
		copy(oa.arr[i+1:], oa.arr[i:])
		oa.arr[i] = item
	} else {
		oa.arr = append(oa.arr, item)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type TweetCounts struct {
	m map[string]*OrderedArray
}

func Constructor() TweetCounts {
	return TweetCounts{map[string]*OrderedArray{}}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	if oa, ok := this.m[tweetName]; ok {
		oa.Add(time)
	} else {
		oa = (&OrderedArray{}).Init([]int{time},
			func(a, b int) bool { return a < b },
			func(a, b int) bool { return a == b },
		)
		this.m[tweetName] = oa
	}
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	o := make([]int, 0, 32)
	var d int
	switch freq {
	case "minute":
		d = 60
	case "hour":
		d = 3600
	case "day":
		d = 86400
	}
	oa := this.m[tweetName]
	b, e := startTime, min(startTime+d, endTime+1)
	for b < e {
		o = append(o, oa.LowerBound(e)-oa.LowerBound(b))
		b += d
		e = min(b+d, endTime+1)
	}
	return o
}
