import "bytes"

var words = []string{"Thousand", "Million", "Billion"}
var words1 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var words2 = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

func numberToWords1k(num int) *bytes.Buffer {
	buf := &bytes.Buffer{}
	a, b, c := num/100, num%100, num%10
	if a > 0 {
		buf.WriteString(words1[a])
		buf.WriteString(" Hundred")
		if b != 0 {
			buf.WriteByte(' ')
		}
	}
	if b >= 20 {
		buf.WriteString(words2[b/10])
		if c != 0 {
			buf.WriteByte(' ')
			buf.WriteString(words1[c])
		}
	} else {
		buf.WriteString(words1[b])
	}
	return buf
}

func numberToWords(num int) string {
	if 0 == num {
		return "Zero"
	}
	buf := numberToWords1k(num % 1000)
	for i := 0; i < 3 && num != 0; i++ {
		num /= 1000
		r := num % 1000
		if r != 0 {
			buf1 := numberToWords1k(r)
			buf1.WriteByte(' ')
			buf1.WriteString(words[i])
			buf1.WriteByte(' ')
			buf.WriteTo(buf1)
			buf = buf1
		}
	}
	bs := buf.Bytes()
	e := len(bs) - 1
	if bs[e] == ' ' {
		bs = bs[:e]
	}
	return string(bs)
}
