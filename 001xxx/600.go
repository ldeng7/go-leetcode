type ThroneInheritance struct {
	k string
	m map[string][]string
	d map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{kingName, map[string][]string{}, map[string]bool{}}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.m[parentName] = append(this.m[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.d[name] = true
}

func (this *ThroneInheritance) cal(name string, ar []string) []string {
	if !this.d[name] {
		ar = append(ar, name)
	}
	for _, n := range this.m[name] {
		ar = this.cal(n, ar)
	}
	return ar
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	return this.cal(this.k, make([]string, 0, 16))
}
