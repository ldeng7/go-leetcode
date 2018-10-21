func intToRoman(num int) string {
	t := []string{"", "M", "MM", "MMM"}
	h := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	d := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	u := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return t[num/1000] + h[(num%1000)/100] + d[(num%100)/10] + u[num%10]
}