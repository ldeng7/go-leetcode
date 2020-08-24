import "fmt"

var m = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6,
	"Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
}

func reformatDate(date string) string {
	day, i := date[0]-'0', 4
	if d := date[1]; d >= '0' && d <= '9' {
		day, i = day*10+(d-'0'), 5
	}
	return fmt.Sprintf("%s-%02d-%02d", date[i+4:], m[date[i:i+3]], day)
}
