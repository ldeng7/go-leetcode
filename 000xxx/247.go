import (
	"reflect"
	"unsafe"
)

var chars3 = [3]byte{'0', '1', '8'}
var chars4a = [4]byte{'1', '6', '8', '9'}
var chars4b = [4]byte{'1', '9', '8', '6'}
var chars5a = [5]byte{'0', '1', '6', '8', '9'}
var chars5b = [5]byte{'0', '1', '9', '8', '6'}

func findStrobogrammatic(n int) []string {
	if 0 == n {
		return []string{""}
	} else if 1 == n {
		return []string{"0", "1", "8"}
	}

	var nStr int
	var bs []byte
	if n&1 == 0 {
		nStr = 4
		for i := 0; i < (n>>1)-1; i++ {
			nStr *= 5
		}
		bs = make([]byte, nStr*n)

		k, m := 0, n>>1
		for i := 0; i < nStr; i++ {
			l := nStr / 4
			ic := i / l
			bs[k], bs[k+n-1] = chars4a[ic], chars4b[ic]
			for j := 1; j < m; j++ {
				l /= 5
				ic = i / l % 5
				bs[k+j], bs[k+n-j-1] = chars5a[ic], chars5b[ic]
			}
			k += n
		}
	} else {
		nStr = 12
		for i := 0; i < (n>>1)-1; i++ {
			nStr *= 5
		}
		bs = make([]byte, nStr*n)

		k, m := 0, n>>1
		for i := 0; i < nStr; i++ {
			l := nStr / 4
			ic := i / l
			bs[k], bs[k+n-1] = chars4a[ic], chars4b[ic]
			for j := 1; j < m; j++ {
				l /= 5
				ic = i / l % 5
				bs[k+j], bs[k+n-j-1] = chars5a[ic], chars5b[ic]
			}
			bs[k+m] = chars3[i%3]
			k += n
		}
	}

	out := make([]string, nStr)
	p := (*reflect.SliceHeader)(unsafe.Pointer(&bs)).Data
	for i := 0; i < nStr; i++ {
		ps := &reflect.StringHeader{p, n}
		out[i] = *(*string)(unsafe.Pointer(ps))
		p += uintptr(n)
	}
	return out
}

