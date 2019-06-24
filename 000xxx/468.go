import (
	"strings"
)

func validIPAddress(IP string) string {
	if -1 == strings.IndexByte(IP, ':') {
		es := strings.Split(IP, ".")
		if len(es) != 4 {
			return "Neither"
		}
		for _, e := range es {
			l := len(e)
			if l == 0 || l > 3 || (l > 1 && e[0] == '0') {
				return "Neither"
			}
			for i := 0; i < l; i++ {
				c := e[i]
				if c < '0' || c > '9' {
					return "Neither"
				}
			}
			if l == 3 && e[0] >= '2' && e[1] >= '5' && e[2] >= '6' {
				return "Neither"
			}
		}
		return "IPv4"
	} else {
		es := strings.Split(IP, ":")
		if len(es) != 8 {
			return "Neither"
		}
		for _, e := range es {
			l := len(e)
			if l == 0 || l > 4 {
				return "Neither"
			}
			for i := 0; i < l; i++ {
				c := e[i]
				if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
}
