package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //remove * at the tail of the path
	compositeEntry := []Entry{}
	//根据后缀名选出JAR文件
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			//返回SkipDir以外的错误，则处理停止
			return err
		}
		if info.IsDir() && path != baseDir {
			//跳过通配符的子目录,返回这个错误只会抛弃这个子目录所有内容而不会报错
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
