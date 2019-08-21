import "time"

func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	return t.YearDay()
}
