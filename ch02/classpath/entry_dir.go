package classpath

import (
	"io/ioutil"
	"path/filepath"
)

/**
目录形式的类路径
*/
type DirEntry struct {
	//绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	//参数转换成绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		//转换过程出现错误，终止程序执行
		panic(err)
	}
	return &DirEntry{absDir}
}

//如果是包含指定类文件的目录，则不报错
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	//将目录和class文件名拼接成完整的路径
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
