import "fmt"

func toHex(num int) string {
	return fmt.Sprintf("%x", uint32(num))
}
