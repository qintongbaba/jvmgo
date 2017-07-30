package classpath

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		errors.New("使用DirEntry读取class信息失败！className:" + className)
	}
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
