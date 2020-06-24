func daysBetweenDates(date1 string, date2 string) int {
	d1, _ := time.Parse("2006-01-02", date1)
	d2, _ := time.Parse("2006-01-02", date2)
	o := int(d1.Sub(d2).Hours()) / 24
	if o < 0 {
		o = -o
	}
	return o
}
