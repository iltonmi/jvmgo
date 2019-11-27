package classpath

import (
	"os"
	"strings"
)

const pathListSeperator = string(os.PathListSeparator) //存放路径分隔符
type Entry interface {
	/**
	寻找和加载class文件
	arg:class文件的相对路径，路径之间用/隔开，文件名必须有class后缀
	return：读取的字节数据、定位到class文件的Entry、错误信息
	*/
	readClass(className string) ([]byte, Entry, error)
	/**
	返回字符串表示
	*/
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeperator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
