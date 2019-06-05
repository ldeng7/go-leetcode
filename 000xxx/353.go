type SnakeGame struct {
	width, height int
	score, iHead  int
	food, snake   [][]int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{width, height, 0, 0, food, [][]int{{0, 0}}}
}

func (this *SnakeGame) Move(direction string) int {
	iTail := this.iHead - 1
	if iTail < 0 {
		iTail = len(this.snake) - 1
	}
	head := this.snake[this.iHead]
	y, x := head[0], head[1]

	switch direction {
	case "U":
		y--
	case "D":
		y++
	case "L":
		x--
	case "R":
		x++
	}
	if y < 0 || y >= this.height || x < 0 || x >= this.width {
		return -1
	}
	for i := 0; i < len(this.snake); i++ {
		if i == iTail {
			continue
		}
		p := this.snake[i]
		if p[0] == y && p[1] == x {
			return -1
		}
	}

	if 0 != len(this.food) && this.food[0][0] == y && this.food[0][1] == x {
		this.snake = append(this.snake, nil)
		copy(this.snake[this.iHead+1:], this.snake[this.iHead:])
		this.snake[this.iHead] = []int{y, x}
		this.food = this.food[1:]
		this.score++
	} else {
		head := this.snake[iTail]
		head[0], head[1] = y, x
		this.iHead = iTail
	}
	return this.score
}
