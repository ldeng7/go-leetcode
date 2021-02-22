type ParkingSystem struct {
	c  [3]int
	ar [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{c: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	carType--
	if this.ar[carType] == this.c[carType] {
		return false
	}
	this.ar[carType]++
	return true
}
