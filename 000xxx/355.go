type pqElemType = [2]int
type pqElemLessCb = func(pqElemType, pqElemType) bool

type PriorityQueue struct {
	arr    []pqElemType
	lessCb pqElemLessCb
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq.lessCb(pq.arr[j], pq.arr[i]) {
			break
		}
		pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
		j = i
	}
}

func (pq *PriorityQueue) down(i0, n int) bool {
	i := i0
	for {
		j1 := i<<1 + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && pq.lessCb(pq.arr[j2], pq.arr[j1]) {
			j = j2
		}
		if !pq.lessCb(pq.arr[j], pq.arr[i]) {
			break
		}
		pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
		i = j
	}
	return i > i0
}

func (pq *PriorityQueue) Init(arr []pqElemType, lessCb pqElemLessCb) *PriorityQueue {
	pq.arr = arr
	pq.lessCb = lessCb
	l := len(pq.arr)
	for i := l>>1 - 1; i >= 0; i-- {
		pq.down(i, l)
	}
	return pq
}

func (pq *PriorityQueue) Len() int {
	return len(pq.arr)
}

func (pq *PriorityQueue) Top() *pqElemType {
	if len(pq.arr) != 0 {
		e := pq.arr[0]
		return &e
	}
	return nil
}

func (pq *PriorityQueue) Push(item pqElemType) {
	pq.arr = append(pq.arr, item)
	pq.up(len(pq.arr) - 1)
}

func (pq *PriorityQueue) Pop() *pqElemType {
	i := len(pq.arr) - 1
	if i >= 0 {
		pq.arr[0], pq.arr[i] = pq.arr[i], pq.arr[0]
		pq.down(0, i)
		e := pq.arr[i]
		pq.arr = pq.arr[:i]
		return &e
	}
	return nil
}

func (pq *PriorityQueue) Set(index int, item pqElemType) {
	pq.arr[index] = item
	if !pq.down(index, len(pq.arr)) {
		pq.up(index)
	}
}

type Twitter struct {
	time    int
	friends map[int]map[int]bool
	tweets  map[int][][2]int
}

func Constructor() Twitter {
	return Twitter{0, map[int]map[int]bool{}, map[int][][2]int{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Follow(userId, userId)
	this.tweets[userId] = append(this.tweets[userId], [2]int{this.time, tweetId})
	this.time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	q := (&PriorityQueue{}).Init(nil, func(a, b [2]int) bool { return a[0] < b[0] })
	for uid, _ := range this.friends[userId] {
		tweets := this.tweets[uid]
		ie := len(tweets) - 10
		if ie < 0 {
			ie = 0
		}
		for i := len(tweets) - 1; i >= ie; i-- {
			p := tweets[i]
			if q.Len() == 10 {
				if (*q.Top())[0] > p[0] {
					break
				}
				q.Set(0, p)
			} else {
				q.Push(p)
			}
		}
	}
	o, i := make([]int, q.Len()), q.Len()-1
	for q.Len() != 0 {
		o[i], i = (*q.Pop())[1], i-1
	}
	return o
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if m, ok := this.friends[followerId]; ok {
		m[followeeId] = true
	} else {
		this.friends[followerId] = map[int]bool{followeeId: true}
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId != followeeId {
		if m, ok := this.friends[followerId]; ok {
			delete(m, followeeId)
		}
	}
}
