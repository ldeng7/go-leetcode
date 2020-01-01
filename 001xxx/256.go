import "fmt"

func encode(num int) string {
	return fmt.Sprintf("%b", num+1)[1:]
}
