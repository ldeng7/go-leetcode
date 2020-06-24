func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type BrowserHistory struct {
	ar []string
	i  int
}

func Constructor(homepage string) BrowserHistory {
	ar := make([]string, 1, 5001)
	ar[0] = homepage
	return BrowserHistory{ar, 1}
}

func (this *BrowserHistory) Visit(url string) {
	if this.i < len(this.ar) {
		this.ar = this.ar[:this.i]
	}
	this.ar = append(this.ar, url)
	this.i++
}

func (this *BrowserHistory) Back(steps int) string {
	this.i = max(this.i-steps, 1)
	return this.ar[this.i-1]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.i = min(this.i+steps, len(this.ar))
	return this.ar[this.i-1]
}
