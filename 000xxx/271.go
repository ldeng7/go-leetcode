import (
	"bytes"
	"strconv"
)

type Codec struct{}

func (codec *Codec) Encode(strs []string) string {
	buf := &bytes.Buffer{}
	for _, s := range strs {
		ls := strconv.Itoa(len(s))
		buf.WriteByte(byte(len(ls)))
		buf.WriteString(ls)
		buf.WriteString(s)
	}
	return buf.String()
}

func (codec *Codec) Decode(strs string) []string {
	out := []string{}
	i := 0
	for i < len(strs) {
		lls := int(strs[i])
		i++
		ls, _ := strconv.Atoi(strs[i : i+lls])
		i += lls
		out = append(out, strs[i:i+ls])
		i += ls
	}
	return out
}
