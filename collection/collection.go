package collection

import (
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/collection", new(Collection))
}

type Collection struct {
	mu    sync.Mutex
	items map[string][]*Item
}

type Item struct {
	FileName   string
	FilePath   string
	FileData   []byte
	FileSize   int64
	MimeType   string
	ParrentDir string
}

func (c *Collection) CreateCollection(collectionPath string) {
	c.items = make(map[string][]*Item)
	err := filepath.WalkDir(collectionPath,
		func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				c.mu.Lock()
				c.items[filepath.Dir(path)] = append(c.items[filepath.Dir(path)], &Item{
					FileName:   filepath.Base(path),
					FilePath:   path,
					ParrentDir: filepath.Dir(path),
				})
				c.mu.Unlock()
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func PrepareItem(item *Item) {
	data, err := os.ReadFile(item.FilePath)
	if err != nil {
		log.Println(err)
	}
	item.FileData = data
	item.MimeType = http.DetectContentType(data)

	file_info, err := os.Stat(item.FilePath)
	if err != nil {
		log.Println(err)
	}
	item.FileSize = file_info.Size()

}

func (c *Collection) GetAllItems() []*Item {
	for _, items := range c.items {
		for _, item := range items {
			PrepareItem(item)
		}
		return items
	}
	return []*Item{}
}

func randMapKey(m map[string][]*Item) string {
	mapKeys := make([]string, 0, len(m))
	for key := range m {
		mapKeys = append(mapKeys, key)
	}
	return mapKeys[rand.Intn(len(mapKeys))]
}

func (c *Collection) GetRandomItem() *Item {
	key := randMapKey(c.items)
	item := c.items[key][rand.Intn(len(c.items[key]))]
	PrepareItem(item)
	return item
}

func (c *Collection) GetItemByFilename(filename string) *Item {
	for _, items := range c.items {
		for _, item := range items {
			if item.FileName == filename {
				PrepareItem(item)
				return item
			}
		}
	}
	return &Item{}
}

func (c *Collection) GetItemByFilepath(filepath string) *Item {
	for _, items := range c.items {
		for _, item := range items {
			if item.FilePath == filepath {
				PrepareItem(item)
				return item
			}
		}
	}
	return &Item{}
}

func (c *Collection) GetItemsByParrentDir(parrentDir string) []*Item {
	for key, items := range c.items {
		if strings.Contains(key, parrentDir) {
			for _, item := range items {
				PrepareItem(item)
			}
			return items
		}
	}
	return []*Item{}
}
