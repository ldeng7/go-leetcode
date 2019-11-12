import "strings"

type FileSystem struct {
	m map[string]int
}

func Constructor() FileSystem {
	return FileSystem{map[string]int{
		"": -1,
	}}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	if _, ok := this.m[path]; ok {
		return false
	}
	if _, ok := this.m[path[:strings.LastIndexByte(path, '/')]]; !ok {
		return false
	}
	this.m[path] = value
	return true
}

func (this *FileSystem) Get(path string) int {
	if v, ok := this.m[path]; ok {
		return v
	}
	return -1
}
