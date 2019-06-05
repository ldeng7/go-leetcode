import (
	"sort"
	"strings"
)

type FileSystem struct {
	dirs  map[string]map[string]bool
	files map[string]string
}

func Constructor() FileSystem {
	return FileSystem{map[string]map[string]bool{}, map[string]string{}}
}

func (this *FileSystem) Ls(path string) []string {
	if _, ok := this.files[path]; ok {
		return []string{path[strings.LastIndexByte(path, '/')+1:]}
	}
	out := []string{}
	for k, _ := range this.dirs[path] {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (this *FileSystem) touchDir(dir string) {
	if _, ok := this.dirs[dir]; !ok {
		this.dirs[dir] = map[string]bool{}
	}
}

func (this *FileSystem) Mkdir(path string) {
	if "/" == path {
		this.touchDir(path)
		return
	}
	es := strings.Split(path, "/")
	for i := 1; i < len(es); i++ {
		dir := "/" + strings.Join(es[1:i], "/")
		this.touchDir(dir)
		this.dirs[dir][es[i]] = true
	}
	this.touchDir(path)
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	i := strings.LastIndexByte(filePath, '/')
	dir, file := filePath[:i], filePath[i+1:]
	if 0 == len(dir) {
		dir = "/"
	}
	if _, ok := this.dirs[dir]; !ok {
		this.Mkdir(dir)
	}

	this.dirs[dir][file] = true
	if c, ok := this.files[filePath]; ok {
		this.files[filePath] = c + content
	} else {
		this.files[filePath] = content
	}
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	return this.files[filePath]
}
