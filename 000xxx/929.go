import "strings"

func numUniqueEmails(emails []string) int {
	set := map[string]bool{}
	for _, e := range emails {
		ia, i := strings.IndexByte(e, '@'), strings.IndexByte(e, '+')
		if i <= 0 {
			i = ia
		}
		bs := []byte{}
		for j := 0; j < i; j++ {
			c := e[j]
			if c != '.' {
				bs = append(bs, c)
			}
		}
		set[string(bs)+e[ia:]] = true
	}
	return len(set)
}
