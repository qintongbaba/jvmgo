package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	rc, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer rc.Close()

	for _, f := range rc.File {
		if f.Name == className {
			frc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer frc.Close()

			data, err := ioutil.ReadAll(frc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)

}
func (self *ZipEntry) String() string {
	return self.absPath
}
