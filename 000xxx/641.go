type MyCircularDeque struct {
	l, c, h, t int
	arr        []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{k, 0, k - 1, 0, make([]int, k)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.c != this.l {
		this.arr[this.h] = value
		this.h--
		if this.h == -1 {
			this.h = this.l - 1
		}
		this.c++
		return true
	}
	return false
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.c != this.l {
		this.arr[this.t] = value
		this.t++
		if this.t == this.l {
			this.t = 0
		}
		this.c++
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.c != 0 {
		this.h++
		if this.h == this.l {
			this.h = 0
		}
		this.c--
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.c != 0 {
		this.t--
		if this.t == -1 {
			this.t = this.l - 1
		}
		this.c--
		return true
	}
	return false
}

func (this *MyCircularDeque) GetFront() int {
	if this.c != 0 {
		i := this.h + 1
		if i == this.l {
			i = 0
		}
		return this.arr[i]
	}
	return -1
}

func (this *MyCircularDeque) GetRear() int {
	if this.c != 0 {
		i := this.t - 1
		if i == -1 {
			i = this.l - 1
		}
		return this.arr[i]
	}
	return -1
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.c == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.c == this.l
}
