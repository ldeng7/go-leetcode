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

	var bs []byte
	var nStr int
	m := n >> 1
	calInner := func(i int) {
		l, k := nStr/4, n*i
		ic := i / l
		bs[k], bs[k+n-1] = chars4a[ic], chars4b[ic]
		for j := 1; j < m; j++ {
			l /= 5
			ic = i / l % 5
			bs[k+j], bs[k+n-j-1] = chars5a[ic], chars5b[ic]
		}
	}
	var cal func()
	if n&1 == 0 {
		nStr = 4
		cal = func() {
			for i := 0; i < nStr; i++ {
				calInner(i)
			}
		}
	} else {
		nStr = 12
		cal = func() {
			for i := 0; i < nStr; i++ {
				calInner(i)
				bs[n*i+m] = chars3[i%3]
			}
		}
	}

	for i := 0; i < (n>>1)-1; i++ {
		nStr *= 5
	}
	bs = make([]byte, nStr*n)
	cal()

	out := make([]string, nStr)
	p := (*reflect.SliceHeader)(unsafe.Pointer(&bs)).Data
	for i := 0; i < nStr; i++ {
		ps := &reflect.StringHeader{p, n}
		out[i] = *(*string)(unsafe.Pointer(ps))
		p += uintptr(n)
	}
	return out
}
