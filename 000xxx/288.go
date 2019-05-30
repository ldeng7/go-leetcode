import "fmt"

type ValidWordAbbr struct {
	m map[string]string
}

func getAbbr(s string) string {
	if len(s) <= 2 {
		return s
	}
	return fmt.Sprintf("%c%d%c", s[0], len(s)-2, s[len(s)-1])
}

func Constructor(dictionary []string) ValidWordAbbr {
	m := map[string]string{}
	for _, d := range dictionary {
		k := getAbbr(d)
		if _, ok := m[k]; ok && m[k] != d {
			m[k] = ""
		} else {
			m[k] = d
		}
	}
	return ValidWordAbbr{m}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	k := getAbbr(word)
	v, ok := this.m[k]
	return !ok || v == word
}
