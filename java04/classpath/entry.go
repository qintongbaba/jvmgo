package classpath

import (
	"path/filepath"
	"strings"
)

const pathListSeparator = string(filepath.ListSeparator)

//classpath的接口
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

//工厂方法
func newEntry(path string) Entry {

	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") ||
		strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
