import "math/rand"

type RandomizedCollection struct {
	nums []int
	m    map[int]map[int]bool
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{[]int{}, map[int]map[int]bool{}}
}

func (this *RandomizedCollection) Insert(val int) bool {
	if s, ok := this.m[val]; ok {
		s[len(this.nums)] = true
	} else {
		this.m[val] = map[int]bool{len(this.nums): true}
	}
	this.nums = append(this.nums, val)
	return 1 == len(this.m[val])
}

func (this *RandomizedCollection) Remove(val int) bool {
	s := this.m[val]
	if 0 == len(s) {
		return false
	}
	var n int
	for n, _ = range s {
		break
	}
	delete(s, n)
	e := len(this.nums) - 1
	if e != n {
		i := this.nums[e]
		this.nums[n] = i
		s = this.m[i]
		delete(s, e)
		s[n] = true
	}
	this.nums = this.nums[:e]
	return true
}

func (this *RandomizedCollection) GetRandom() int {
	if 0 == len(this.nums) {
		return -1
	}
	return this.nums[rand.Intn(len(this.nums))]
}
