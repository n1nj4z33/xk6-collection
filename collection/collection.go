package collection

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/collection", New())
}

type File struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Data     []byte `json:"data"`
	MimeType string `json:"mime_type"`
	Size     int64  `json:"size"`
}

type Collection struct {
	Name  string                 `json:"name"`
	Dirs  map[string]*Collection `json:"dirs"`
	Files []*File                `json:"files"`
}

func New() *Collection {
	return &Collection{
		Dirs:  map[string]*Collection{},
		Files: []*File{},
	}
}

func (c *Collection) Init(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	c.Name = filepath.Base(path)
	dirInfo, err := f.ReadDir(0)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range dirInfo {
		if info.IsDir() {
			subDir := c.addDir(info.Name())
			subDir.Init(fmt.Sprintf("%s/%s", path, info.Name()))
		} else {
			c.addFile(info.Name(), path)
		}
	}
}

func (c *Collection) addFile(name string, path string) {
	c.Files = append(c.Files, &File{
		Name: name,
		Path: filepath.Join(path, name),
	})
}

func (c *Collection) addDir(name string) *Collection {
	c.Dirs[name] = &Collection{
		Name:  name,
		Dirs:  map[string]*Collection{},
		Files: []*File{},
	}
	return c.Dirs[name]
}

func (c *Collection) PrepareFile(file *File) *File {
	data, err := os.ReadFile(file.Path)
	if err != nil {
		log.Fatal(err)
	}
	file.Data = data
	file.MimeType = http.DetectContentType(data)
	fileInfo, err := os.Stat(file.Path)
	if err != nil {
		log.Fatal(err)
	}
	file.Size = fileInfo.Size()
	return file
}
