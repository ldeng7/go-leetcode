var t = [4]string{"", "M", "MM", "MMM"}
var h = [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var d = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var u = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func intToRoman(num int) string {
	return t[num/1000] + h[(num%1000)/100] + d[(num%100)/10] + u[num%10]
}
