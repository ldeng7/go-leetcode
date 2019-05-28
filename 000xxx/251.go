type Vector2D struct {
	v        [][]int
	i, j, ei int
}

func Constructor(v [][]int) Vector2D {
	i, j := 0, 0
	ei := len(v) - 1
	for ; i <= ei; i++ {
		if 0 != len(v[i]) {
			break
		}
	}
	return Vector2D{v, i, j, ei}
}

func (this *Vector2D) Next() int {
	e := this.v[this.i][this.j]
	this.j++
	for ; this.i <= this.ei; this.i++ {
		if this.j != len(this.v[this.i]) {
			break
		} else {
			this.j = 0
		}
	}
	return e
}

func (this *Vector2D) HasNext() bool {
	return this.i < this.ei || (this.i == this.ei && this.j != len(this.v[this.ei]))
}
