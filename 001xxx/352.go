type ProductOfNumbers struct {
	ar []int
}

func Constructor() ProductOfNumbers {
	ar := make([]int, 1, 16)
	ar[0] = 1
	return ProductOfNumbers{ar}
}

func (this *ProductOfNumbers) Add(num int) {
	if 0 != num {
		ar := this.ar
		this.ar = append(ar, ar[len(ar)-1]*num)
	} else {
		ar := make([]int, 1, 16)
		ar[0] = 1
		this.ar = ar
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	ar := this.ar
	l := len(ar)
	if l > k {
		return ar[l-1] / ar[l-k-1]
	}
	return 0
}
