import (
	"strconv"
	"strings"
)

func nextClosestTime(time string) string {
	bs := []byte{'0', '0', '0', '0'}
	tmp := [4]int{600, 60, 10, 1}
	i := strings.IndexByte(time, ':')
	t, _ := strconv.Atoi(time[:i])
	t1, _ := strconv.Atoi(time[i+1:])
	t = t*60 + t1
	for i = 1; i <= 1440; i++ {
		tn := (t + i) % 1440
		j := 0
		for ; j < 4; j++ {
			bs[j] = '0' + byte(tn/tmp[j])
			tn %= tmp[j]
			if strings.IndexByte(time, bs[j]) == -1 {
				break
			}
		}
		if j == 4 {
			break
		}
	}
	bs = append(bs, bs[3])
	bs[2], bs[3] = ':', bs[2]
	return string(bs)
}
