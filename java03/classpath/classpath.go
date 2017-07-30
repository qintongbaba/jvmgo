package classpath

import (
	"errors"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	classpath := &Classpath{}
	classpath.parseBootAndExtClasspath(jreOption)
	classpath.parseUserClasspath(cpOption)
	return classpath
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	extLibPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(extLibPath)

}
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, error) {
	className = className + ".class"
	if data, _, err := self.bootClasspath.readClass(className); err == nil {
		return data, nil
	}
	if data, _, err := self.extClasspath.readClass(className); err == nil {
		return data, nil
	}
	if data, _, err := self.userClasspath.readClass(className); err == nil {
		return data, nil
	}

	return nil, errors.New("classpath加载不到类信息:" + className)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
