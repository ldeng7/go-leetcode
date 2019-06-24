import (
	"math"
	"math/rand"
)

type Solution struct {
	r, x, y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	theta := math.Pi * rand.Float64() * 2
	l := math.Sqrt(rand.Float64()) * this.r
	return []float64{this.x + l*math.Cos(theta), this.y + l*math.Sin(theta)}
}
